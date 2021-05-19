# LambdaAuthorizerGolang
A simple implementation of Lambda Authorizer in Golang for API gateway. 

What I'm trying to do is I've got this API migrated to AWS and I need a authorizer to protect the api from anonymous access. So the most straight forward way to that is simply write my own customized lambda authorizer and bind to my API gateway. 

I've got this dynamoDB table "ApiUser" to store my client information, to put it simple, it has only a key - ClientID and an attribute - ClientSecret. Both ClientID and ClientSecret are sent to the client and will be used to authenticate the API request.

Each time a client tried to invoke my API, 2 headers and on query string is mandatory:

Query String:
  <p>request_id: <i>this should be a unique string for each request from a client, duplicate request_id from a client should be ignored</i></p>

Header:
  <p>client_id: <i>the ClientID they hold</i></p>
  <p>signature: <i>how to calculate the signature will be explained later</i></p>
  
How to calculate the signature:
1. combine ClientID, ClientSecret and RequestID together in a string like this:
  client_id=aaa&client_secret=bbb&request_id=ccc
2. Hash the string using MD5 (feel free to use other hash methods) and that's how the signature is generated.

Because only clients has the ClientSecret, no one else should be able to work out the signature.

In my authorizer lambda function, I will tried to 
    1. retrieve the client_id from header and request_id from query string
    2. retrieve the ClientSecret from dynamodb 
    3. workout the signature myself 
    4. compare my signature and the one in request's header
    5. if the 2 signatures are identical - pass; throw 403 or 500 otherwise.

AWS recently just upgraded the authorizer payload format to 2.0 and the aws-sdk-go to v2, so there aren't many documentations aronud. So I spend quite some time on this demo and hopefully this repo could help you if you are looking for a similar solution.

<b>Again, this is just a demo, but at least it workable, and feel free to enhanced it and log your thoughts. </b>