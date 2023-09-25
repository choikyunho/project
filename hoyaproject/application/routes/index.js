var express = require('express');
var router = express.Router();
var cc= require('../public/javascripts/cc')
var ccc= require('../public/javascripts/ccc')


/* GET home page. */
router.get('/', async function(req, res, next) {

const usercookie = req.cookies["USER"]
console.log(usercookie)


if(!usercookie){
  res.render('users',{title:'Express'});
}else{
 const userdata = JSON.parse(usercookie)
////////////////////////////////////

CCresult = await cc.cc_call(userdata.userid,"GetInputmoney","")
CCCresult = await ccc.ccc_call(userdata.userid,"GetOutputmoney","")
console.log(CCresult)
console.log(CCCresult)
////////////////////////////////////

//화면에 랜더링은 1번만 된다.
  res.render('index', {title:'Express',username:userdata.name, result:CCresult,result1:CCCresult});
}

});

router.get('/logout', function(req, res, next){
  res.clearCookie("USER")
  res.redirect('/')
});

module.exports = router;
