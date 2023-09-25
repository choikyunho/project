var express = require('express');
var router = express.Router();
var cc= require('../public/javascripts/cc')

/* GET home page. */
router.get('/', async function(req, res, next) {
res.render('inputmoney',{title:'Express'})
});

router.post('/', async function(req, res, next) {
    const inputmoney = req.body.inputmoney
    const usercookie = req.cookies["USER"]
    const userdata = JSON.parse(usercookie)
    CCresult = await cc.cc_call(userdata.userid, "Inputmoney",inputmoney)
    console.log(CCresult)
    res.redirect("/")
    });

module.exports = router;
