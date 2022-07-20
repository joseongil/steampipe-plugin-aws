select
  arn,
  name,
  description,
  firewall_policy_status,
  region,
  tags
from
  aws.aws_networkfirewall_firewall_policy
where 
  akas = '["{{ output.resource_aka.value }}"]';
