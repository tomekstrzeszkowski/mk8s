from django.urls import include, path
from rest_framework import routers

from app.views import LeadViewSet

router = routers.DefaultRouter()
router.register(r'leads', LeadViewSet)

urlpatterns = [
    path('', include(router.urls)),
]