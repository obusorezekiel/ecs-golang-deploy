variable "dbname" {
  type = string
}

variable "dbuser" {
  type      = string
  sensitive = true
}

variable "dbpass" {
  type      = string
  sensitive = true
}

