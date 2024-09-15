variable "lb_sg" {}
variable "public_subnets" {
    type = list(string)
}
variable "tg_port" {}
variable "tg_protocol" {}
variable "vpc_id" {}
variable "listener_port" {}
variable "listener_protocol" {}