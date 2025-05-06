/*
Copyright 2025.

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

package tkmoperator

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/types"
)

/*
func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()

	err := CreateCsiDriverObject(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Printf("Closing Connection\n")
	}()

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		var count uint64
		for range ticker.C {
			count = count + 1
			log.Printf("Still Alive: %d\n", count)
		}
	}()

	<-stop

	log.Printf("Exiting...\n")
}
*/

const (
	CSI_DRIVER_YAML = "./config/bpfman-operator/csi-driver-info.yaml"
	CSI_DRIVER_NAME = "csi.tkm.io"
)

func (r *TritonKernelCacheReconciler) CreateCsiDriverObject(ctx context.Context) error {
	tkmCsiDriver := &storagev1.CSIDriver{}

	// one-shot try to create TKM's CSIDriver object if it doesn't exist, does not re-trigger reconcile.
	if err := r.Get(ctx, types.NamespacedName{Namespace: corev1.NamespaceAll, Name: CSI_DRIVER_NAME}, tkmCsiDriver); err != nil {
		if errors.IsNotFound(err) {
			tkmCsiDriver = LoadCsiDriverFromFile(CSI_DRIVER_YAML)

			log.Printf("Calling KubeAPI to Create TKM CSI Driver object")
			if err := r.Create(ctx, tkmCsiDriver); err != nil {
				r.Logger.Error(err, "Failed to create Bpfman csi driver")
				return ctrl.Result{Requeue: true, RequeueAfter: retryDurationOperator}, nil
			}
		} else {
			r.Logger.Error(err, "Failed to get csi.bpfman.io csidriver")
		}
	}
}

func LoadCsiDriverFromFile(path string) *storagev1.CSIDriver {
	// Load static CSIDriver yaml from disk
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, _ := decode(b, nil, nil)

	return obj.(*storagev1.CSIDriver)
}
