select akas, name, region, tags, title
from aws.aws_elastic_beanstalk_application
where name = '{{ resourceName }}';