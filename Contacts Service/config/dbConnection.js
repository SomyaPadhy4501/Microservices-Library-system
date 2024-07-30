const mongoose = require('mongoose');
const dotenv = require("dotenv").config();

const connectDb = async () => {
    try {
     const connect = await mongoose.connect(process.env.CONNECTION_STRING);
     console.log(
        "Connected to the database" , connect.connection.name
     );
    } catch (err) {
        console.log(err);
    }
}

module.exports = connectDb;