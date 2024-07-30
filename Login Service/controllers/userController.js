const asyncHandler = require("express-async-handler");
const User = require("../models/usersModel");
const jwt = require("jsonwebtoken")
const bcryptjs = require("bcryptjs")

//GET "/api/users/current" get all users

const currentUsers = asyncHandler(async(req,res) => {
    const user = await User.find();
   res.status(200).json(req.user);
});

//POST "/api/registerusers" create new user
//@access PUBLIC

const registerUsers = asyncHandler(async(req,res) => {
    const{name , email , password} = req.body;
    if(!name||!email||!password){
        res.status(400);
        throw new Error ("All fields mandatory");
    }
    const Availableuser = await User.findOne({email});
    if(Availableuser) {
        res.status(400);
        throw new Error ("User already registered");
    }
    //Hash password
    const hashedPassword = await bcryptjs.hash(password , 10);
    const user = await User.create({
        name,
        email,
        password : hashedPassword,
    });
    console.log("User created successfully");
    if(user)
        {
            res.status(201).json({ _id:user.id , email:user.email});
        } else {
            res.status(400);
            throw new Error("User invalid");
        }
 });

 //CURRENT "/api/users/login" update exsisting user 
 //@access PUBLIC

 const loginUsers = asyncHandler(async(req,res) => {
    const {email , password} = req.body;
    if(!email||!password){
        res.status(400);
        throw new Error("Enter email and password");
    }
    const user = await User.findOne({email});
    if(user && (await bcryptjs.compare(password , user.password))){
        const accessToken = jwt.sign({
            user:{
                name: user.name,
                email: user.email,
                id: user.id,
            },
        }, process.env.ACCESS_TOKEN_SECERET , 
    {expiresIn: "10m"});
        res.status(200).json({accessToken});
    } else{
        res.status(401)
        throw new Error("Email or Password not valid");
    }
 });


module.exports = {currentUsers , registerUsers , loginUsers} ; 