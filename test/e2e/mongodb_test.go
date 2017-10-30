package e2e_test

import (
	"github.com/appscode/go/hold"
	"github.com/appscode/go/types"
	tapi "github.com/k8sdb/apimachinery/apis/kubedb/v1alpha1"
	"github.com/k8sdb/mongodb/test/e2e/framework"
	"github.com/k8sdb/mongodb/test/e2e/matcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	//"github.com/the-redback/go-oneliners"
	//"fmt"
	"fmt"

	"github.com/the-redback/go-oneliners"
)

const (
	S3_BUCKET_NAME       = "S3_BUCKET_NAME"
	GCS_BUCKET_NAME      = "GCS_BUCKET_NAME"
	AZURE_CONTAINER_NAME = "AZURE_CONTAINER_NAME"
	SWIFT_CONTAINER_NAME = "SWIFT_CONTAINER_NAME"
)

var _ = Describe("MongoDB", func() {
	var (
		err      error
		f        *framework.Invocation
		mongodb  *tapi.MongoDB
		snapshot *tapi.Snapshot
		//secret      *core.Secret
		skipMessage string
	)

	BeforeEach(func() {
		f = root.Invoke()
		mongodb = f.MongoDB()
		snapshot = f.Snapshot()
		skipMessage = ""
	})

	var createAndWaitForRunning = func() {
		By("Create MongoDB: " + mongodb.Name)
		err = f.CreateMongoDB(mongodb)
		Expect(err).NotTo(HaveOccurred())

		By("Wait for Running mongodb")
		f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())
	}

	var deleteTestResource = func() {
		By("Delete mongodb")
		err = f.DeleteMongoDB(mongodb.ObjectMeta)
		Expect(err).NotTo(HaveOccurred())

		By("Wait for mongodb to be paused")
		f.EventuallyDormantDatabaseStatus(mongodb.ObjectMeta).Should(matcher.HavePaused())

		By("WipeOut mongodb")
		_, err := f.TryPatchDormantDatabase(mongodb.ObjectMeta, func(in *tapi.DormantDatabase) *tapi.DormantDatabase {
			in.Spec.WipeOut = true
			return in
		})
		Expect(err).NotTo(HaveOccurred())

		By("Wait for mongodb to be wipedOut")
		f.EventuallyDormantDatabaseStatus(mongodb.ObjectMeta).Should(matcher.HaveWipedOut())

		err = f.DeleteDormantDatabase(mongodb.ObjectMeta)
		Expect(err).NotTo(HaveOccurred())
	}

	var shouldSuccessfullyRunning = func() {
		if skipMessage != "" {
			Skip(skipMessage)
		}

		// Create MongoDB
		createAndWaitForRunning()

		hold.Hold()

		// Delete test resource
		deleteTestResource()
	}

	Describe("Test", func() {

		Context("General", func() {

			Context("-", func() {
				It("should run successfully", shouldSuccessfullyRunning)
			})

			Context("With PVC", func() {
				BeforeEach(func() {
					if f.StorageClass == "" {
						skipMessage = "Missing StorageClassName. Provide as flag to test this."
					}
					mongodb.Spec.Storage = &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("5Gi"),
							},
						},
						StorageClassName: types.StringP(f.StorageClass),
					}
				})
				It("should run successfully", shouldSuccessfullyRunning)
			})
		})

		Context("DoNotPause", func() {
			BeforeEach(func() {
				mongodb.Spec.DoNotPause = true
			})

			It("should work successfully", func() {
				// Create and wait for running MongoDB
				createAndWaitForRunning()

				By("Delete mongodb")
				err = f.DeleteMongoDB(mongodb.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("MongoDB is not paused. Check for mongodb")
				f.EventuallyMongoDB(mongodb.ObjectMeta).Should(BeTrue())

				By("Check for Running mongodb")
				f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())

				By("Update mongodb to set DoNotPause=false")
				f.TryPatchMongoDB(mongodb.ObjectMeta, func(in *tapi.MongoDB) *tapi.MongoDB {
					in.Spec.DoNotPause = false
					return in
				})

				// Delete test resource
				deleteTestResource()
			})
		})

		//Context("Snapshot", func() {
		//	var skipDataCheck bool
		//
		//	AfterEach(func() {
		//		f.DeleteSecret(secret.ObjectMeta)
		//	})
		//
		//	BeforeEach(func() {
		//		skipDataCheck = false
		//		snapshot.Spec.DatabaseName = mongodb.Name
		//	})
		//
		//	var shouldTakeSnapshot = func() {
		//		// Create and wait for running MongoDB
		//		createAndWaitForRunning()
		//
		//		By("Create Secret")
		//		f.CreateSecret(secret)
		//
		//		By("Create Snapshot")
		//		f.CreateSnapshot(snapshot)
		//
		//		By("Check for Successed snapshot")
		//		f.EventuallySnapshotPhase(snapshot.ObjectMeta).Should(Equal(tapi.SnapshotPhaseSuccessed))
		//
		//		if !skipDataCheck {
		//			By("Check for snapshot data")
		//			f.EventuallySnapshotDataFound(snapshot).Should(BeTrue())
		//		}
		//
		//		// Delete test resource
		//		deleteTestResource()
		//
		//		if !skipDataCheck {
		//			By("Check for snapshot data")
		//			f.EventuallySnapshotDataFound(snapshot).Should(BeFalse())
		//		}
		//	}
		//
		//	Context("In Local", func() {
		//		BeforeEach(func() {
		//			skipDataCheck = true
		//			secret = f.SecretForLocalBackend()
		//			snapshot.Spec.StorageSecretName = secret.Name
		//			snapshot.Spec.Local = &tapi.LocalSpec{
		//				Path: "/repo",
		//				VolumeSource: core.VolumeSource{
		//					HostPath: &core.HostPathVolumeSource{
		//						Path: "/repo",
		//					},
		//				},
		//			}
		//		})
		//
		//		It("should take Snapshot successfully", shouldTakeSnapshot)
		//	})
		//
		//	Context("In S3", func() {
		//		BeforeEach(func() {
		//			secret = f.SecretForS3Backend()
		//			snapshot.Spec.StorageSecretName = secret.Name
		//			snapshot.Spec.S3 = &tapi.S3Spec{
		//				Bucket: os.Getenv(S3_BUCKET_NAME),
		//			}
		//		})
		//
		//		It("should take Snapshot successfully", shouldTakeSnapshot)
		//	})
		//
		//	Context("In GCS", func() {
		//		BeforeEach(func() {
		//			secret = f.SecretForGCSBackend()
		//			snapshot.Spec.StorageSecretName = secret.Name
		//			snapshot.Spec.GCS = &tapi.GCSSpec{
		//				Bucket: os.Getenv(GCS_BUCKET_NAME),
		//			}
		//		})
		//
		//		It("should take Snapshot successfully", shouldTakeSnapshot)
		//	})
		//
		//	Context("In Azure", func() {
		//		BeforeEach(func() {
		//			secret = f.SecretForAzureBackend()
		//			snapshot.Spec.StorageSecretName = secret.Name
		//			snapshot.Spec.Azure = &tapi.AzureSpec{
		//				Container: os.Getenv(AZURE_CONTAINER_NAME),
		//			}
		//		})
		//
		//		It("should take Snapshot successfully", shouldTakeSnapshot)
		//	})
		//
		//	Context("In Swift", func() {
		//		BeforeEach(func() {
		//			secret = f.SecretForSwiftBackend()
		//			snapshot.Spec.StorageSecretName = secret.Name
		//			snapshot.Spec.Swift = &tapi.SwiftSpec{
		//				Container: os.Getenv(SWIFT_CONTAINER_NAME),
		//			}
		//		})
		//
		//		It("should take Snapshot successfully", shouldTakeSnapshot)
		//	})
		//})

		Context("Initialize", func() {
			Context("With Script", func() {
				BeforeEach(func() {
					mongodb.Spec.Init = &tapi.InitSpec{
						ScriptSource: &tapi.ScriptSourceSpec{
							VolumeSource: core.VolumeSource{
								GitRepo: &core.GitRepoVolumeSource{
									Repository: "https://github.com/the-redback/k8s-mongodb-init-script.git",
									Directory:  ".",
								},
							},
						},
					}
				})

				It("should run successfully", shouldSuccessfullyRunning)

			})

			//Context("With Snapshot", func() {
			//	AfterEach(func() {
			//		f.DeleteSecret(secret.ObjectMeta)
			//	})
			//
			//	BeforeEach(func() {
			//		secret = f.SecretForS3Backend()
			//		snapshot.Spec.StorageSecretName = secret.Name
			//		snapshot.Spec.S3 = &tapi.S3Spec{
			//			Bucket: os.Getenv(S3_BUCKET_NAME),
			//		}
			//		snapshot.Spec.DatabaseName = mongodb.Name
			//	})
			//
			//	It("should run successfully", func() {
			//		// Create and wait for running MongoDB
			//		createAndWaitForRunning()
			//
			//		By("Create Secret")
			//		f.CreateSecret(secret)
			//
			//		By("Create Snapshot")
			//		f.CreateSnapshot(snapshot)
			//
			//		By("Check for Successed snapshot")
			//		f.EventuallySnapshotPhase(snapshot.ObjectMeta).Should(Equal(tapi.SnapshotPhaseSuccessed))
			//
			//		By("Check for snapshot data")
			//		f.EventuallySnapshotDataFound(snapshot).Should(BeTrue())
			//
			//		oldMongoDB, err := f.GetMongoDB(mongodb.ObjectMeta)
			//		Expect(err).NotTo(HaveOccurred())
			//
			//		By("Create mongodb from snapshot")
			//		mongodb = f.MongoDB()
			//		mongodb.Spec.DatabaseSecret = oldMongoDB.Spec.DatabaseSecret
			//		mongodb.Spec.Init = &tapi.InitSpec{
			//			SnapshotSource: &tapi.SnapshotSourceSpec{
			//				Namespace: snapshot.Namespace,
			//				Name:      snapshot.Name,
			//			},
			//		}
			//
			//		// Create and wait for running MongoDB
			//		createAndWaitForRunning()
			//
			//		// Delete test resource
			//		deleteTestResource()
			//		mongodb = oldMongoDB
			//		// Delete test resource
			//		deleteTestResource()
			//	})
			//})
		})

		Context("Resume", func() {
			var usedInitSpec bool
			BeforeEach(func() {
				usedInitSpec = false
			})

			var shouldResumeSuccessfully = func() {
				// Create and wait for running MongoDB
				createAndWaitForRunning()

				By("Delete mongodb")
				f.DeleteMongoDB(mongodb.ObjectMeta)

				By("Wait for mongodb to be paused")
				f.EventuallyDormantDatabaseStatus(mongodb.ObjectMeta).Should(matcher.HavePaused())

				_, err = f.TryPatchDormantDatabase(mongodb.ObjectMeta, func(in *tapi.DormantDatabase) *tapi.DormantDatabase {
					in.Spec.Resume = true
					return in
				})
				Expect(err).NotTo(HaveOccurred())

				By("Wait for DormantDatabase to be deleted")
				f.EventuallyDormantDatabase(mongodb.ObjectMeta).Should(BeFalse())

				By("Wait for Running mongodb")
				f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())

				mongodb, err = f.GetMongoDB(mongodb.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				if usedInitSpec {
					Expect(mongodb.Spec.Init).Should(BeNil())
					Expect(mongodb.Annotations[tapi.MongoDBInitSpec]).ShouldNot(BeEmpty())
				}

				// Delete test resource
				deleteTestResource()
			}

			Context("-", func() {
				It("should resume DormantDatabase successfully", shouldResumeSuccessfully)
			})

			Context("With Init", func() {
				BeforeEach(func() {
					usedInitSpec = true
					mongodb.Spec.Init = &tapi.InitSpec{
						ScriptSource: &tapi.ScriptSourceSpec{
							VolumeSource: core.VolumeSource{
								GitRepo: &core.GitRepoVolumeSource{
									Repository: "https://github.com/the-redback/k8s-mongodb-init-script.git",
									Directory:  ".",
								},
							},
						},
					}
				})

				It("should resume DormantDatabase successfully", shouldResumeSuccessfully)
			})

			Context("With original MongoDB", func() {
				It("should resume DormantDatabase successfully", func() {
					// Create and wait for running MongoDB
					createAndWaitForRunning()
					By("Delete mongodb")
					f.DeleteMongoDB(mongodb.ObjectMeta)

					By("Wait for mongodb to be paused")
					f.EventuallyDormantDatabaseStatus(mongodb.ObjectMeta).Should(matcher.HavePaused())

					// Create MongoDB object again to resume it
					By("Create MongoDB: " + mongodb.Name)
					err = f.CreateMongoDB(mongodb)
					Expect(err).NotTo(HaveOccurred())

					By("Wait for DormantDatabase to be deleted")
					f.EventuallyDormantDatabase(mongodb.ObjectMeta).Should(BeFalse())

					By("Wait for Running mongodb")
					f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())

					_, err = f.GetMongoDB(mongodb.ObjectMeta)
					Expect(err).NotTo(HaveOccurred())

					if usedInitSpec {
						Expect(mongodb.Spec.Init).Should(BeNil())
						Expect(mongodb.Annotations[tapi.MongoDBInitSpec]).ShouldNot(BeEmpty())
					}

					// Delete test resource
					deleteTestResource()
				})

				Context("with init", func() {
					BeforeEach(func() {
						usedInitSpec = true
						mongodb.Spec.Init = &tapi.InitSpec{
							ScriptSource: &tapi.ScriptSourceSpec{
								VolumeSource: core.VolumeSource{
									GitRepo: &core.GitRepoVolumeSource{
										Repository: "https://github.com/the-redback/k8s-mongodb-init-script.git",
										Directory:  ".",
									},
								},
							},
						}
					})

					FIt("should resume DormantDatabase successfully", func() {
						// Create and wait for running MongoDB
						createAndWaitForRunning()

						for i := 0; i < 3; i++ {
							By(">>>>>>>>>>>>>> "+fmt.Sprintf("%v", i) + " times running <<<<<<<<<<<")
							By("Delete mongodb")
							f.DeleteMongoDB(mongodb.ObjectMeta)

							By("Wait for mongodb to be paused")
							f.EventuallyDormantDatabaseStatus(mongodb.ObjectMeta).Should(matcher.HavePaused())

							// Create MongoDB object again to resume it
							By("Create MongoDB: " + mongodb.Name)
							err = f.CreateMongoDB(mongodb)
							if err != nil {
								oneliners.FILE(err)
							}
							Expect(err).NotTo(HaveOccurred())

							By("Wait for DormantDatabase to be deleted")
							f.EventuallyDormantDatabase(mongodb.ObjectMeta).Should(BeFalse())

							By("Wait for Running mongodb")
							f.EventuallyMongoDBRunning(mongodb.ObjectMeta).Should(BeTrue())

							_mongodb, err := f.GetMongoDB(mongodb.ObjectMeta)
							Expect(err).NotTo(HaveOccurred())
							oneliners.PrettyJson(_mongodb, "new mongo")
						}

						// Delete test resource
						deleteTestResource()
					})
				})
			})
		})

		//Context("SnapshotScheduler", func() {
		//	AfterEach(func() {
		//		f.DeleteSecret(secret.ObjectMeta)
		//	})
		//
		//	BeforeEach(func() {
		//		secret = f.SecretForLocalBackend()
		//	})
		//
		//	Context("With Startup", func() {
		//		BeforeEach(func() {
		//			mongodb.Spec.BackupSchedule = &tapi.BackupScheduleSpec{
		//				CronExpression: "@every 1m",
		//				SnapshotStorageSpec: tapi.SnapshotStorageSpec{
		//					StorageSecretName: secret.Name,
		//					Local: &tapi.LocalSpec{
		//						Path: "/repo",
		//						VolumeSource: core.VolumeSource{
		//							HostPath: &core.HostPathVolumeSource{
		//								Path: "/repo",
		//							},
		//						},
		//					},
		//				},
		//			}
		//		})
		//
		//		It("should run schedular successfully", func() {
		//			By("Create Secret")
		//			f.CreateSecret(secret)
		//
		//			// Create and wait for running MongoDB
		//			createAndWaitForRunning()
		//
		//			By("Count multiple Snapshot")
		//			f.EventuallySnapshotCount(mongodb.ObjectMeta).Should(matcher.MoreThan(3))
		//
		//			deleteTestResource()
		//		})
		//	})
		//
		//	Context("With Update", func() {
		//		It("should run schedular successfully", func() {
		//			// Create and wait for running MongoDB
		//			createAndWaitForRunning()
		//
		//			By("Create Secret")
		//			f.CreateSecret(secret)
		//
		//			By("Update mongodb")
		//			_, err = f.TryPatchMongoDB(mongodb.ObjectMeta, func(in *tapi.MongoDB) *tapi.MongoDB {
		//				in.Spec.BackupSchedule = &tapi.BackupScheduleSpec{
		//					CronExpression: "@every 1m",
		//					SnapshotStorageSpec: tapi.SnapshotStorageSpec{
		//						StorageSecretName: secret.Name,
		//						Local: &tapi.LocalSpec{
		//							Path: "/repo",
		//							VolumeSource: core.VolumeSource{
		//								HostPath: &core.HostPathVolumeSource{
		//									Path: "/repo",
		//								},
		//							},
		//						},
		//					},
		//				}
		//
		//				return in
		//			})
		//			Expect(err).NotTo(HaveOccurred())
		//
		//			By("Count multiple Snapshot")
		//			f.EventuallySnapshotCount(mongodb.ObjectMeta).Should(matcher.MoreThan(3))
		//
		//			deleteTestResource()
		//		})
		//	})
		//})

	})
})
