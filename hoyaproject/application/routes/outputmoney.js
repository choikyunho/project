var express = require('express');
var router = express.Router();
var ccc= require('../public/javascripts/ccc')

/* GET home page. */
router.get('/', async function(req, res, next) {
res.render('outputmoney',{title:'Express'})
});

router.post('/', async function(req, res, next) {
    const outputmoney = req.body.outputmoney
    const outputmoney1 = req.body.purpose

    var args = [outputmoney,outputmoney1]

    const usercookie = req.cookies["USER"]
    const userdata = JSON.parse(usercookie)

    CCCresult = await ccc.ccc_call(userdata.userid, "Outputmoney",args)
    console.log(CCCresult)
    res.redirect("/")
    });

module.exports = router;
