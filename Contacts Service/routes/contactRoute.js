const express = require("express");
const router = express.Router();
const {getContacts, getContact , deleteContacts , updateContacts , createContacts} = require("../controllers/contactController");
const valid = require("../middleware/validateToken");

router.use(valid);
router.route('/').get(getContacts).post(createContacts);
router.route('/:id').get(getContact).put(updateContacts).delete(deleteContacts);

module.exports = router;