resource "aws_lb" "complex_lb" {
  name            = "complex-lb"
  security_groups = [var.lb_sg]
  subnets         = var.public_subnets
  idle_timeout    = 400
}

resource "aws_lb_target_group" "complex_tg" {
  name     = "complex-lb-tg-${substr(uuid(), 0, 8)}"
  port     = var.tg_port
  protocol = var.tg_protocol
  vpc_id   = var.vpc_id
  target_type = "ip"

  lifecycle {
    ignore_changes        = [name]
    create_before_destroy = true
  }
}

resource "aws_lb_listener" "complex_lb_listener" {
  load_balancer_arn = aws_lb.complex_lb.arn
  port              = var.listener_port
  protocol          = var.listener_protocol

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.complex_tg.arn
  }
}
