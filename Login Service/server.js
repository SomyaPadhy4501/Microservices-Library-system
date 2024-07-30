const express = require('express');
const errorHandler = require('./middleware/errorHandler');
const connectDb = require('./config/dbConnection');
const dotenv = require("dotenv").config();
const app = express()
const port = process.env.PORT || 4000; 

connectDb();
app.use(express.json());
app.use("/api/users" , require("./routes/userRoute"));
app.use(errorHandler)


app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})

