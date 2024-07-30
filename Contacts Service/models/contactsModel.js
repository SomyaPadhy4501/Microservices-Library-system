const mongoose = require('mongoose');


const contactSchema = mongoose.Schema({
   user_id: {
      type: mongoose.Schema.Types.ObjectId,
      required: true,
      reference: "User",
   },
 name:{
    type:String,
    required: [true, "Please enter a name"]
 },
 email:{
    type:String,
    required: [true, "Enter an email"]
 },
 phone: {
    type: Number,
    required:[true , "Enter a phone number"]
 },
 address: {
    type: String,
    required:[true , "Enter a phone number"]
 },
}, {
    timestamps: true
});

module.exports = mongoose.model("Contact" , contactSchema);