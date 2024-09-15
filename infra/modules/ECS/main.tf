resource "aws_ecs_cluster" "golang_ecs_cluster" {
  name = var.ecs_cluster
}

resource "aws_ecs_task_definition" "golang_task_definition" {
  family                   = var.task_family
  container_definitions    = <<DEFINITION
  [
    {
      "name": "my_go_task",
      "image": "353928175117.dkr.ecr.us-east-1.amazonaws.com/gin_gorm_crud:latest",
      "essential": true,
      "portMappings": [
        {
          "containerPort": 8888,
          "hostPort": 8888
        }
      ],
      "memory": 512,
      "cpu": 256
    }
  ]
  DEFINITION
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = 512    
  cpu                      = 256
  execution_role_arn       = aws_iam_role.ecsTaskExecutionRole.arn
}

resource "aws_iam_role" "ecsTaskExecutionRole" {
  name               = "ecsTaskExecutionRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
}

data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "ecsTaskExecutionRole_policy" {
  role       = aws_iam_role.ecsTaskExecutionRole.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_ecs_service" "app_service" {
  name            = "my_ecs_cluster_service"
  cluster         = aws_ecs_cluster.golang_ecs_cluster.id
  task_definition = aws_ecs_task_definition.golang_task_definition.arn
  launch_type     = "FARGATE"
  desired_count   = 3

  load_balancer {
    target_group_arn = var.tg_arn
    container_name   = "my_go_task"
    container_port   = 8888
  }

  network_configuration {
    subnets         = var.private_subnets  # Updated syntax
    security_groups = [var.ecs_sg]  # Ensure var.ecs_sg is a list
  }
}
