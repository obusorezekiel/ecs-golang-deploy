output "alb_dns" {
  value = aws_lb.complex_lb.dns_name
}

output "lb_endpoint" {
  value = aws_lb.complex_lb.dns_name
}

output "lb_tg_name" {
  value = aws_lb_target_group.complex_tg.name
}

output "lb_tg" {
  value = aws_lb_target_group.complex_tg.arn
}