# Rhyvu

## Summary

Rhyvu is a CMS framework with basic functionality for blogs and pages. Configuration is possible through an admin module, and the file-based data store ([Storm](https://github.com/asdine/storm)) requires no user configuration or installation.

## Install and Use

Not yet available.

## Comparison 

If you need a flexible community favorite, I recommend [Ponzu](https://github.com/ponzu-cms/ponzu), which is a solid, production ready headless CMS with an automatic JSON API, featuring auto-HTTPS from Let's Encrypt, HTTP/2 Server Push, and a flexible server framework.

Rhyvu differs in that (besides being in pre-alpha) it is strongly opinionated. Generating a site with Rhyvu will result in a frontend with blog and page modules running on Bootstrap 4, JQuery 3, and HTML with an admin UI to customize menus, features and settings. The backend uses Go for server logic and [Storm](https://github.com/asdine/storm) to store data, but if you aren't designing a new module, this shouldn't be relevant.

## Roadmap

There is not yet much usable functionality. I've grabbed some code from my misc and dwn projects, but I haven't started cleaning it up and putting it together yet. The following features are prioritized:

- Auth
- Pages module
- Admin module
- Blog module

## Developing

Developing has additional dependencies. You will need Git, Make, and Go. Soon I might add NPM.