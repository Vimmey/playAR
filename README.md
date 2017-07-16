Project playAR

Unique features : 

1. First Screen : Augmented reality 
               A. Lat,Long : Camera-direction-wise : Marker/Km/Title/Rating
               B. Menu On Click : Items , Cost Of Two

2. Intermediate Screen: Delivery/Pickup

				A. Pickup : Eta Based On Abs(Traffic - Restaurent Preparing Time)
                B. Profit : No Swiggy Resources/Delivery Person Used

3. Tracking Screen : Game 
				A. To Engage Customers: Tic-tac-toe : Till Delivery                 
                
                B. Other Games :
                    Create Combos : If It Becomes Trend, Some Credits
                    Snake : Longer The Snake
                  	Quiz : How Well U Know Ur City Restaurant

                C. Game Credits Conversion To Swiggy Money For Future Orders

API's exposed :

/restaurants
/menu


Request Format : 

POST Request to : /restaurants

#<post-params>

{
"latitude":"",
"longitude":"",


}

Response Format : 


[{
"id":"",
"latitude":"",
"longitude":"",
"altitude":"",
"title":"",
"description":"",
"distance":"",
"rating":"",
"pickupTime":"",
"costForTwo":""
}]

Example : 
[
{
"id":"363222332",
"latitude":"17.934847",
"longitude":"75.616123",
"altitude":"",
"title":"rest_363222332",
"description":"",
"distance":"2.34",
"rating":"",
"pickupTime":"",
"costForTwo":""
},
{
"id":"363915580",
"latitude":"15.459199",
"longitude":"83.385102",
"altitude":"",
"title":"rest_363915580",
"description":"",
"distance":"2.34",
"rating":"",
"pickupTime":"",
"costForTwo":""
},
{
"id":"364256036",
"latitude":"17.934672",
"longitude":"75.616922",
"altitude":"",
"title":"rest_364256036",
"description":"",
"distance":"2.34",
"rating":"",
"pickupTime":"",
"costForTwo":""
}
]


POST Request to : /menu

#<post-params>

{
"id":""




}

Response Format : 

{
"menu_details":[
{
"item_name":"Pepperoni Supremo Pizza ",
"item_price":"1350.00"
},
{
"item_name":"Spicy Medi Pizza ",
"item_price":"490.00"
},
{
"item_name":"Indiano Rustica Pizza ",
"item_price":"1070.00"
},
{
"item_name":"Frijole Pizza ",
"item_price":"690.00"
},
{
"item_name":"Chicken Spicy Cilantro Pizza ",
"item_price":"1010.00"
},
{
"item_name":"Chicken Fajita Pizza ",
"item_price":"1230.00"
},
{
"item_name":"Meat Cravers Pizza ",
"item_price":"1450.00"
},
{
"item_name":"Chipotle Chicken Pizza ",
"item_price":"1190.00"
},
{
"item_name":"Roasted Garlic Chicken Pizza ",
"item_price":"1190.00"
},
{
"item_name":"Fanta",
"item_price":"70.00"
},
{
"item_name":"California Kung Pao Spaghetti Pasta",
"item_price":"850.00"
},
{
"item_name":"Original BBQ Chicken Sandwich",
"item_price":"850.00"
},
{
"item_name":"Chocolate Trinity",
"item_price":"590.00"
},
{
"item_name":"Tomato Basil Spaghetti Pasta",
"item_price":"950.00"
},
{
"item_name":"Red Velvet",
"item_price":"630.00"
}
]
}

