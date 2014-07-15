'use strict';

angular.module('moby')
.controller('MainCtrl', function ($scope, $famous) {
    $scope.footerGridOptions = {
           dimensions: [4,1], // specifies number of columns and rows
    };
    var nav_footer_image_size = [undefined,50];
    $scope.navs = [{class: "footer_nav_home", size: nav_footer_image_size}, {class: "footer_nav_popular", size: nav_footer_image_size}, {class: "footer_nav_wish_list", size: nav_footer_image_size}, {class: "footer_nav_cart", size: nav_footer_image_size}];
});
