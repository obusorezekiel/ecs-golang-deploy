variable "ecs_cluster" {
  description = "The name of the ECS cluster"
  type        = string
}

variable "task_family" {
  description = "The ECS task definition family"
  type        = string
}


variable "ecs_sg" {
  description = "Security group ID for ECS tasks"
  type        = string
}

variable "vpc_id" {
  description = "VPC ID for the ECS service and task"
  type        = string
}

variable "listener_port" {
  description = "Port for the load balancer listener"
  type        = number

}

variable "listener_protocol" {
  description = "Protocol for the load balancer listener"
  type        = string
}

variable "tg_arn" {
  type = string
}


variable "tg_port" {
  description = "Port for the target group"
  type        = number
}

variable "tg_protocol" {
  description = "Protocol for the target group"
  type        = string
}

variable "private_subnets" {
  type = list(string)
}
