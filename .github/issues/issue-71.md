---
code: null
id: MDU6SXNzdWUyOTU2MDc4MTQ=
databaseId: 295607814
number: 71
url: https://github.com/octopot/forma/issues/71
title: "admin ui"
issueType: null
labels:
  - "status:blocked"
  - "status:research"
assignees:
  - kamilsk
milestone: "[[milestone-4]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-02-08T17:52:45Z
updatedAt: 2018-10-17T09:42:48Z
lastEditedAt: 2018-03-01T09:03:02Z
closedAt: 2018-10-17T09:42:48Z
---

# admin ui

check:
- https://www.creative-tim.com/product/material-dashboard-react
- https://www.creative-tim.com/product/vue-light-bootstrap-dashboard
- https://www.creative-tim.com/product/paper-kit-2-angular
- https://www.creative-tim.com/product/light-bootstrap-dashboard-react
- https://www.creative-tim.com/product/now-ui-dashboard

more here https://github.com/creativetimofficial

<!-- 2018-03-30T08:11Z https://github.com/octopot/forma/issues/71#issuecomment-377474043
https://tabler.github.io looks good
-->

<!-- 2018-04-01T17:05Z https://github.com/octopot/forma/issues/71#issuecomment-377800931
for the quick start, it can be a good choice: https://laravelvoyager.com

nginx' config file for php app can look like this
```
server {
    listen 80 default_server;

    root /var/www/html/public;

    index index.html index.htm index.php;

    server_name _;

    charset utf-8;

    location = /favicon.ico { log_not_found off; access_log off; }
    location = /robots.txt  { log_not_found off; access_log off; }

    location / {
        try_files $uri $uri/ /index.php$is_args$args;
    }

    location ~ \.php$ {
        include snippets/fastcgi-php.conf;
        fastcgi_pass php:9000;
    }

    error_page 404 /index.php;

    location ~ /\.ht {
        deny all;
    }
}
```
-->

<!-- 2018-04-13T09:02Z https://github.com/octopot/forma/issues/71#issuecomment-381072698
will be a part of client package
-->

<!-- 2018-06-07T18:27Z https://github.com/octopot/forma/issues/71#issuecomment-395519371
https://github.com/flatlogic/sing-app looks good
-->
