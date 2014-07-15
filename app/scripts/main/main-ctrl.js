'use strict';

angular.module('moby')
.controller('MainCtrl', function ($scope, $famous) {
    $scope.footerGridOptions = {
           dimensions: [4,1], // specifies number of columns and rows
    };
    var nav_footer_image_size = [undefined,50];
    $scope.navs = [{class: "footer_nav_home", size: nav_footer_image_size}, {class: "footer_nav_search", size: nav_footer_image_size}, {class: "footer_nav_popular", size: nav_footer_image_size}, {class: "footer_nav_cart", size: nav_footer_image_size}];



    var EventHandler = $famous['famous/core/EventHandler'];
    $scope.eventHandler = new EventHandler();

    $scope.list = [{content: "famous"}, {content: "angular"}, {content: "rocks!"},{content: "famous"}, {content: "angular"}, {content: "rocks!"},{content: "famous"}, {content: "angular"}, {content: "rocks!"},{content: "famous"}, {content: "angular"}, {content: "rocks!"}];

    $scope.options = {
          myScrollView: {
        paginated: false,
        speedLimit: 5,
        direction: 1,
      }
    };
});
