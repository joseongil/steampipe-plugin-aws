select name, arn, partition, region
from aws.aws_sagemaker_app
where user_profile_name = '{{ output.user_profile_name.value }}';