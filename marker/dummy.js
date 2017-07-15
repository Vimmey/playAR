(function(){
  'use strict';

  angular
    .module('app')   
    .factory('dummyService', function($resource){
      var data = $resource('../api/public/user', {     
       get: {
          method:'POST'
        }  
      });
      return data;
    });
})();
