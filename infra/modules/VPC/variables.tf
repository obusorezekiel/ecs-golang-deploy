variable "vpc_cidr" {
    type = string
}

variable "public_sn_count" {
    type = number
}

variable "private_sn_count" {
  type = number
}

variable "db_subnet_group" {
  type = bool
}
