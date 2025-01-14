/*
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License.
 */

const { REDIRECT_URI } = require('../authConfig');
const authProvider = require('../auth/AuthProvider');

var express = require('express');
var router = express.Router();

router.get('/', function (req, res, next) {
    res.render('index', {
        title: 'MSAL Node & Express Web App',
        isAuthenticated: req.session.isAuthenticated,
        username: req.session.account?.username,
    });
});

    

router.post('/getAToken', authProvider.handleRedirect({
    successRedirect: '/',
    failureRedirect: '/auth/signin'
}));


module.exports = router;
