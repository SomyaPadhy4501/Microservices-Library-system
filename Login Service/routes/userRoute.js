const express = require("express");
const router = express.Router();
const {currentUsers, loginUsers, registerUsers} = require("../controllers/userController");
const valid = require ("../middleware/validateToken");

router.post("/login" , loginUsers).post("/register" , registerUsers);
router.get("/current" , valid ,currentUsers);

module.exports = router;