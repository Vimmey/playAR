(function(){
  'use strict';

  const myaccModal = {
    templateUrl: 'views/myaccount.modal.html',
    controller: myaccCtrl,
    controllerAs: 'myacc'
  };
  angular
    .module('app')
    .component('myaccModal', myaccModal);

  myaccCtrl.$inject = ['$element', '$attrs', 'UserService', 'Upload', '$timeout'];

  function myaccCtrl($element, $attrs, UserService, Upload, $timeout){
    var vm = this;