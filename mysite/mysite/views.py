from django.http import HttpResponse

def index(request):
    return HttpResponse("hello world ,your are in the polls index.")