resource "google_compute_network" "private_network" {
  name = "test-vpc"

  # defaults to true.  false = --subnet-mode custom
  auto_create_subnetworks = "false"
}

resource "google_compute_subnetwork" "kubernetes" {
  name                     = "kubernetes-subnet"
  ip_cidr_range            = "10.0.0.0/24"
  network                  = google_compute_network.private_network.id
  private_ip_google_access = "true"
  region                   = local.region
}
