resource "aws_ecr_repository" "gin_gorm_crud" {
  name = "gin_gorm_crud"
}


locals {
  instance_type = "t2.small"
  location      = "us-east-1"
  environment   = "dev"
  vpc_cidr      = "10.123.0.0/16"
}

module "vpc" {
  source = "./modules/VPC"
  vpc_cidr = local.vpc_cidr
  public_sn_count = 2
  private_sn_count = 2
  db_subnet_group = true

}

module "sg" {
  source = "./modules/SG"
  vpc_id = module.vpc.vpc_id
  lb_sg_name = "complex_lb_sg"
  backend_sg_name = "complex_backend_ecs_sg"
  rds_sg_name = "complex_rds_sg"
  http_port = 80
  https_port = 443
  postgres_port = 5432
  ingress_cidr_blocks = ["0.0.0.0/0"]
  egress_cidr_blocks = ["0.0.0.0/0"]
}

module "ALB" {
  source = "./modules/ALB"
  lb_sg = module.sg.lb_security_group_id
  public_subnets = module.vpc.public_subnet_ids
  tg_port = 80
  tg_protocol = "HTTP"
  vpc_id = module.vpc.vpc_id
  listener_port     = 80
  listener_protocol = "HTTP"
}

module "ECS" {
  source = "./modules/ECS"
  ecs_cluster = "golang_ecs_cluster"
  private_subnets = module.vpc.private_subnet_ids
  task_family = "my_go_task"
  tg_arn = module.ALB.lb_tg
  ecs_sg = module.sg.backend_security_group_id
  vpc_id = module.vpc.vpc_id
  listener_port = 80
  listener_protocol = "HTTP"
  tg_port = 80
  tg_protocol = "HTTP"
}

module "RDS" {
  source = "./modules/RDS"
  db_storage = 10
  db_engine_version = "12.15"
  db_instance_class = "db.t3.micro"
  db_name = var.dbname
  dbuser = var.dbuser
  dbpassword = var.dbpass
  db_subnet_group_name = module.vpc.db_subnet_group_name
  db_identifier = "complex-db"
  skip_db_snapshot = true
  rds_sg = module.sg.rds_security_group_id
}
