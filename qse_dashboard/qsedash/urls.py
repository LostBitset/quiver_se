# Django URLconf

from django.urls import path

from . import views

urlpatterns = [
    path("", views.index, name="index"),
    path("live.json", views.live_json, name="live_json"),
]

