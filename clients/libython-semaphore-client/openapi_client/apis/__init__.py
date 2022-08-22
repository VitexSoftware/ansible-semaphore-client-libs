
# flake8: noqa

# Import all APIs into this package.
# If you have many APIs here with many many models used in each API this may
# raise a `RecursionError`.
# In order to avoid this, import only the API that you directly need like:
#
#   from openapi_client.api.authentication_api import AuthenticationApi
#
# or import this package, but before doing it, use:
#
#   import sys
#   sys.setrecursionlimit(n)

# Import APIs into API package:
from openapi_client.api.authentication_api import AuthenticationApi
from openapi_client.api.default_api import DefaultApi
from openapi_client.api.project_api import ProjectApi
from openapi_client.api.projects_api import ProjectsApi
from openapi_client.api.schedule_api import ScheduleApi
from openapi_client.api.user_api import UserApi
