resource "aws_db_instance" "complex_db" {
  allocated_storage      = var.db_storage
  engine                 = "postgres"
  engine_version         = var.db_engine_version
  instance_class         = var.db_instance_class
  db_name                = var.db_name
  username               = var.dbuser
  password               = var.dbpassword
  db_subnet_group_name   = var.db_subnet_group_name
  identifier             = var.db_identifier
  skip_final_snapshot    = var.skip_db_snapshot
  vpc_security_group_ids = var.rds_sg

  tags = {
    Name = "complex-db"
  }
}
