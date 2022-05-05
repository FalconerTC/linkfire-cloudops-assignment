resource "google_container_cluster" "kubernetes" {
  provider = google-beta

  name               = "k8s-cluster"
  location           = local.zone
  initial_node_count = 1

  node_config {
    machine_type = "e2-medium"
    preemptible  = true
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }

  network    = google_compute_network.private_network.id
  subnetwork = google_compute_subnetwork.kubernetes.id

  release_channel {
    channel = "REGULAR"
  }

}
