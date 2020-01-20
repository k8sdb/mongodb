/*
Copyright The KubeDB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package e2e_test

import (
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	"kubedb.dev/mongodb/test/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
)

var _ = Describe("MongoDB SSL", func() {

	var (
		err             error
		f               *framework.Invocation
		mongodb         *api.MongoDB
		garbageMongoDB  *api.MongoDBList
		snapshotPVC     *core.PersistentVolumeClaim
		secret          *core.Secret
		skipMessage     string
		verifySharding  bool
		enableSharding  bool
		dbName          string
		clusterAuthMode *api.ClusterAuthMode
		sslMode         *api.SSLMode
	)

	BeforeEach(func() {
		f = root.Invoke()
		mongodb = f.MongoDBStandalone()
		garbageMongoDB = new(api.MongoDBList)
		secret = nil
		skipMessage = ""
		verifySharding = false
		enableSharding = false
		dbName = "kubedb"
		clusterAuthMode = nil
		sslMode = nil
	})

	AfterEach(func() {
		// Cleanup
		By("Cleanup Left Overs")
		By("Delete left over MongoDB objects")
		root.CleanMongoDB()
		By("Delete left over workloads if exists any")
		root.CleanWorkloadLeftOvers()

		if snapshotPVC != nil {
			err := f.DeletePersistentVolumeClaim(snapshotPVC.ObjectMeta)
			if err != nil && !kerr.IsNotFound(err) {
				Expect(err).NotTo(HaveOccurred())
			}
		}
	})

	JustAfterEach(func() {
		if CurrentGinkgoTestDescription().Failed {
			f.PrintDebugHelpers()
		}
	})

	var createAndWaitForRunning = func() {
		By("Create MongoDB: " + mongodb.Name)
		err = f.CreateMongoDB(mongodb)
		Expect(err).NotTo(HaveOccurred())

		By("Wait for Running mongodb")
		f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())

		By("Wait for AppBinding to create")
		f.EventuallyAppBinding(mongodb.ObjectMeta).Should(BeTrue())

		By("Check valid AppBinding Specs")
		err := f.CheckAppBindingSpec(mongodb.ObjectMeta)
		Expect(err).NotTo(HaveOccurred())

		By("Ping mongodb database")
		f.EventuallyPingMongo(mongodb.ObjectMeta)
	}

	var deleteTestResource = func() {
		if mongodb == nil {
			Skip("Skipping")
		}

		By("Check if mongodb " + mongodb.Name + " exists.")
		mg, err := f.GetMongoDB(mongodb.ObjectMeta)
		if err != nil && kerr.IsNotFound(err) {
			// MongoDB was not created. Hence, rest of cleanup is not necessary.
			return
		}
		Expect(err).NotTo(HaveOccurred())

		By("Update mongodb to set spec.terminationPolicy = WipeOut")
		_, err = f.PatchMongoDB(mg.ObjectMeta, func(in *api.MongoDB) *api.MongoDB {
			in.Spec.TerminationPolicy = api.TerminationPolicyWipeOut
			return in
		})
		Expect(err).NotTo(HaveOccurred())

		By("Delete mongodb")
		err = f.DeleteMongoDB(mongodb.ObjectMeta)
		if err != nil && kerr.IsNotFound(err) {
			// MongoDB was not created. Hence, rest of cleanup is not necessary.
			return
		}
		Expect(err).NotTo(HaveOccurred())

		By("Wait for mongodb to be deleted")
		f.EventuallyMongoDB(mongodb.ObjectMeta).Should(BeFalse())

		By("Wait for mongodb resources to be wipedOut")
		f.EventuallyWipedOut(mongodb.ObjectMeta).Should(Succeed())
	}

	var shouldRunWithPVC = func() {
		if skipMessage != "" {
			Skip(skipMessage)
		}
		// Create MongoDB
		createAndWaitForRunning()

		By("Checking SSL settings (if enabled any)")
		f.EventuallyUserSSLSettings(mongodb.ObjectMeta, clusterAuthMode, sslMode).Should(BeTrue())

		if enableSharding {
			By("Enable sharding for db:" + dbName)
			f.EventuallyEnableSharding(mongodb.ObjectMeta, dbName).Should(BeTrue())
		}
		if verifySharding {
			By("Check if db " + dbName + " is set to partitioned")
			f.EventuallyCollectionPartitioned(mongodb.ObjectMeta, dbName).Should(Equal(enableSharding))
		}

		By("Insert Document Inside DB")
		f.EventuallyInsertDocument(mongodb.ObjectMeta, dbName, 3).Should(BeTrue())

		By("Checking Inserted Document")
		f.EventuallyDocumentExists(mongodb.ObjectMeta, dbName, 3).Should(BeTrue())

		By("Delete mongodb")
		err = f.DeleteMongoDB(mongodb.ObjectMeta)
		Expect(err).NotTo(HaveOccurred())

		By("Wait for mongodb to be paused")
		f.EventuallyMongoDB(mongodb.ObjectMeta).Should(BeFalse())

		// Create MongoDB object again to resume it
		By("Create MongoDB: " + mongodb.Name)
		err = f.CreateMongoDB(mongodb)
		Expect(err).NotTo(HaveOccurred())

		By("Wait for Running mongodb")
		f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())

		By("Ping mongodb database")
		f.EventuallyPingMongo(mongodb.ObjectMeta)

		if verifySharding {
			By("Check if db " + dbName + " is set to partitioned")
			f.EventuallyCollectionPartitioned(mongodb.ObjectMeta, dbName).Should(Equal(enableSharding))
		}

		By("Checking Inserted Document")
		f.EventuallyDocumentExists(mongodb.ObjectMeta, dbName, 3).Should(BeTrue())
	}

	var shouldFailToCreateDB = func() {
		By("Create MongoDB: " + mongodb.Name)
		err = f.CreateMongoDB(mongodb)
		Expect(err).To(HaveOccurred())
	}

	var shouldInitializeFromScript = func() {
		// Create and wait for running MongoDB
		createAndWaitForRunning()

		By("Checking SSL settings (if enabled any)")
		f.EventuallyUserSSLSettings(mongodb.ObjectMeta, clusterAuthMode, sslMode).Should(BeTrue())

		if enableSharding {
			By("Enable sharding for db:" + dbName)
			f.EventuallyEnableSharding(mongodb.ObjectMeta, dbName).Should(BeTrue())
		}
		if verifySharding {
			By("Check if db " + dbName + " is set to partitioned")
			f.EventuallyCollectionPartitioned(mongodb.ObjectMeta, dbName).Should(Equal(enableSharding))
		}

		By("Checking Inserted Document from initialization part")
		f.EventuallyDocumentExists(mongodb.ObjectMeta, dbName, 1).Should(BeTrue())

		By("Insert Document Inside DB")
		f.EventuallyInsertDocument(mongodb.ObjectMeta, dbName, 50).Should(BeTrue())

		By("Checking Inserted Document")
		f.EventuallyDocumentExists(mongodb.ObjectMeta, dbName, 50).Should(BeTrue())
	}

	Describe("Test", func() {

		BeforeEach(func() {
			if f.StorageClass == "" {
				Skip("Missing StorageClassName. Provide as flag to test this.")
			}
		})

		AfterEach(func() {
			// Delete test resource
			deleteTestResource()

			for _, mg := range garbageMongoDB.Items {
				*mongodb = mg
				// Delete test resource
				deleteTestResource()
			}

			if secret != nil {
				err := f.DeleteSecret(secret.ObjectMeta)
				if !kerr.IsNotFound(err) {
					Expect(err).NotTo(HaveOccurred())
				}
			}
		})

		Context("General SSL", func() {

			Context("With sslMode requireSSL", func() {

				BeforeEach(func() {
					sslMode = framework.SSLModeP(api.SSLModeRequireSSL)
				})

				Context("Standalone", func() {
					BeforeEach(func() {
						mongodb = f.MongoDBStandalone()
						mongodb.Spec.SSLMode = *sslMode
					})

					It("should run successfully", shouldRunWithPVC)

					// Snapshot doesn't work yet for requireSSL SSLMode
				})

				Context("With ClusterAuthMode keyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})
				})

				Context("With ClusterAuthMode x509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeX509)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})
				})

				Context("With ClusterAuthMode sendkeyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})
				})

				Context("With ClusterAuthMode sendX509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendX509)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})
				})
			})

			Context("With sslMode preferSSL", func() {

				BeforeEach(func() {
					sslMode = framework.SSLModeP(api.SSLModePreferSSL)
				})

				Context("Standalone", func() {

					BeforeEach(func() {
						mongodb = f.MongoDBStandalone()
						mongodb.Spec.SSLMode = *sslMode
					})

					It("should run successfully", shouldRunWithPVC)

					Context("Initialization - script & snapshot", func() {
						var configMap *core.ConfigMap

						BeforeEach(func() {
							configMap = f.ConfigMapForInitialization()
							err := f.CreateConfigMap(configMap)
							Expect(err).NotTo(HaveOccurred())
						})

						AfterEach(func() {
							err := f.DeleteConfigMap(configMap.ObjectMeta)
							Expect(err).NotTo(HaveOccurred())
						})

						BeforeEach(func() {
							secret = f.SecretForGCSBackend()
							mongodb.Spec.Init = &api.InitSpec{
								ScriptSource: &api.ScriptSourceSpec{
									VolumeSource: core.VolumeSource{
										ConfigMap: &core.ConfigMapVolumeSource{
											LocalObjectReference: core.LocalObjectReference{
												Name: configMap.Name,
											},
										},
									},
								},
							}
						})

						It("should run successfully", shouldInitializeFromScript)
					})
				})

				Context("With ClusterAuthMode keyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())
							})

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})

				Context("With ClusterAuthMode x509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeX509)
					})

					Context("With Replica Set", func() {
						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})

				Context("With ClusterAuthMode sendkeyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())
							})

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})

				Context("With ClusterAuthMode sendX509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendX509)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())
							})

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})
			})

			Context("With sslMode allowssl", func() {

				BeforeEach(func() {
					sslMode = framework.SSLModeP(api.SSLModeAllowSSL)
				})

				Context("Standalone", func() {

					BeforeEach(func() {
						mongodb = f.MongoDBStandalone()
						mongodb.Spec.SSLMode = *sslMode
					})

					It("should run successfully", shouldRunWithPVC)

					Context("Initialization - script & snapshot", func() {
						var configMap *core.ConfigMap

						BeforeEach(func() {
							configMap = f.ConfigMapForInitialization()
							err := f.CreateConfigMap(configMap)
							Expect(err).NotTo(HaveOccurred())
						})

						AfterEach(func() {
							err := f.DeleteConfigMap(configMap.ObjectMeta)
							Expect(err).NotTo(HaveOccurred())
						})
						BeforeEach(func() {
							secret = f.SecretForGCSBackend()
							mongodb.Spec.Init = &api.InitSpec{
								ScriptSource: &api.ScriptSourceSpec{
									VolumeSource: core.VolumeSource{
										ConfigMap: &core.ConfigMapVolumeSource{
											LocalObjectReference: core.LocalObjectReference{
												Name: configMap.Name,
											},
										},
									},
								},
							}
						})

						It("should run successfully", shouldInitializeFromScript)
					})
				})

				Context("With ClusterAuthMode keyFile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())
							})

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})

				// should fail. error: BadValue: cannot have x.509 cluster authentication in allowSSL mode
				Context("With ClusterAuthMode x509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeX509)
					})

					Context("With Replica Set", func() {
						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})
				})

				Context("With ClusterAuthMode sendkeyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())
							})

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})

				//should fail. error: BadValue: cannot have x.509 cluster authentication in allowSSL mode
				Context("With ClusterAuthMode sendX509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendX509)
					})

					Context("With Replica Set", func() {
						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})
				})

			})

			Context("With sslMode disabled", func() {

				BeforeEach(func() {
					sslMode = framework.SSLModeP(api.SSLModeDisabled)
				})

				Context("Standalone", func() {
					BeforeEach(func() {
						mongodb = f.MongoDBStandalone()
						mongodb.Spec.SSLMode = *sslMode
					})

					It("should run successfully", shouldRunWithPVC)

					Context("Initialization - script & snapshot", func() {
						var configMap *core.ConfigMap

						AfterEach(func() {
							err := f.DeleteConfigMap(configMap.ObjectMeta)
							Expect(err).NotTo(HaveOccurred())
						})

						BeforeEach(func() {
							configMap = f.ConfigMapForInitialization()
							err := f.CreateConfigMap(configMap)
							Expect(err).NotTo(HaveOccurred())

							secret = f.SecretForGCSBackend()
							mongodb.Spec.Init = &api.InitSpec{
								ScriptSource: &api.ScriptSourceSpec{
									VolumeSource: core.VolumeSource{
										ConfigMap: &core.ConfigMapVolumeSource{
											LocalObjectReference: core.LocalObjectReference{
												Name: configMap.Name,
											},
										},
									},
								},
							}

							mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
						})

						It("should run successfully", shouldInitializeFromScript)
					})
				})

				Context("With ClusterAuthMode keyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeKeyFile)
					})

					Context("With Replica Set", func() {

						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should run successfully", shouldRunWithPVC)

						Context("Initialization - script & snapshot", func() {
							var configMap *core.ConfigMap

							AfterEach(func() {
								err := f.DeleteConfigMap(configMap.ObjectMeta)
								Expect(err).NotTo(HaveOccurred())
							})

							BeforeEach(func() {
								configMap = f.ConfigMapForInitialization()
								err := f.CreateConfigMap(configMap)
								Expect(err).NotTo(HaveOccurred())

								verifySharding = true
								enableSharding = true
								secret = f.SecretForGCSBackend()
								mongodb.Spec.Init = &api.InitSpec{
									ScriptSource: &api.ScriptSourceSpec{
										VolumeSource: core.VolumeSource{
											ConfigMap: &core.ConfigMapVolumeSource{
												LocalObjectReference: core.LocalObjectReference{
													Name: configMap.Name,
												},
											},
										},
									},
								}

								mongodb = f.MongoDBWithFlexibleProbeTimeout(mongodb)
							})

							It("should initialize database successfully", shouldInitializeFromScript)
						})
					})
				})

				// should fail. error: BadValue: need to enable SSL via the sslMode flag
				Context("With ClusterAuthMode x509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeX509)
					})

					Context("With Replica Set", func() {
						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})
				})

				// should fail. error: BadValue: need to enable SSL via the sslMode flag
				Context("With ClusterAuthMode sendkeyfile", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendKeyFile)
					})

					Context("With Replica Set", func() {
						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})

					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})
				})

				// should fail. error: need to enable SSL via the sslMode flag
				Context("With ClusterAuthMode sendX509", func() {

					BeforeEach(func() {
						clusterAuthMode = framework.ClusterAuthModeP(api.ClusterAuthModeSendX509)
					})

					// should fail. error: need to enable SSL via the sslMode flag
					Context("With Replica Set", func() {
						BeforeEach(func() {
							mongodb = f.MongoDBRS()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})

					// should fail. error: need to enable SSL via the sslMode flag
					Context("With Sharding", func() {

						BeforeEach(func() {
							verifySharding = true
							enableSharding = true

							mongodb = f.MongoDBShard()
							mongodb.Spec.ClusterAuthMode = *clusterAuthMode
							mongodb.Spec.SSLMode = *sslMode
						})

						It("should fail creating mongodb object", shouldFailToCreateDB)
					})
				})
			})
		})
	})
})
