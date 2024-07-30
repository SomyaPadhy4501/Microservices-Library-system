const asyncHandler = require("express-async-handler");
const Contact = require("../models/contactsModel");
//Get all contacts
//route is GET /api/contacts
//access private

const getContacts = asyncHandler( async(req , res) => {
    const contacts = await Contact.find({user_id: req.user.id});
    res.status(200).json(contacts);
});

//Create contacts
//route is POST /api/contacts
//access private

const createContacts = asyncHandler(async(req , res) => {
    const {name , email , phone , address} = req.body;
    if(!name || !email || !phone || !address) {
        res.status(400);
        throw new Error("All fields needed");
    }
    const contacts = await Contact.create({
        name,
        email,
        phone,
        address,
        user_id:req.user.id,
    });
    res.status(201).json(contacts);
});

//route is GET /api/contacts/:id
//access private
const getContact = asyncHandler( async(req , res) => {
    const contacts = await Contact.findById(req.params.id);
    if(!contacts){
        res.status(404);
        throw new Error("Contact not found");
    }
    res.status(200).json(contacts);
});

//route is PUT /api/contacts/:id
//access private

const updateContacts = asyncHandler( async(req , res) => {
    const contacts = await Contact.findById(req.params.id);
    if(!contacts){
        res.status(404);
        throw new Error("Contact not found");
    }
    if(contacts.user_id.toString() !== req.user.id){
        res.status(403);
        throw new Error("No perms to update others contacts");
    }
    const updateContact = await Contact.findByIdAndUpdate(
        req.params.id,
        req.body,
        {new:true}
    );
    res.status(200).json(updateContact);
});

//route is DELETE /api/contacts/:id
//access private

const deleteContacts = asyncHandler( async(req , res) => {
    const contacts = await Contact.findById(req.params.id);
    if(!contacts){
        res.status(404);
        throw new Error("Contact not found");
    }
    if(contacts.user_id.toString() !== req.user.id){
        res.status(403);
        throw new Error("No perms to delete others contacts");
    }
    await Contact.deleteOne( {_id: req.params.id});
    res.status(200).json(contacts);
});

module.exports = {getContacts , createContacts , getContact , updateContacts , deleteContacts};