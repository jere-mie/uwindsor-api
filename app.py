from flask import Flask, request, jsonify, make_response
import json
from difflib import SequenceMatcher
from flask_cors import CORS, cross_origin

def similar(a,b): # returns how similar two strings are
    return SequenceMatcher(None, a, b).ratio()

# this is for getting the secret key
with open('secrets.json') as f:
    data = json.load(f)

# creating an instance of Flask as our app
app = Flask(__name__)
app.config['SECRET_KEY'] = data['secret_key']
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

@app.route('/api/<choice>', methods=['GET'])
@cross_origin()
def api(choice):
    keys = {
        "buildings":"building_name",
        "courses":"Course Code",
        "emergency_services":"service",
        "holidays":"date",
        "staff":"Name"
    }

    if choice not in keys: # if the route is invalid, for example /api/rooch
        return make_response(jsonify(f'/api/{choice} is not a valid route'), 400)

    data = "" # getting json from file
    with open(f'static/{choice}.json', 'r') as f:
        data = json.loads(f.read())

    if 'q' not in request.args: # we return all data if no param
        return make_response(jsonify(data), 200)

    choice_key = keys[choice] # getting the name of the attribute to query by in the json
    key = 0 # the key of the data found
    max = -1 # max similarity between query and each json item

    # finding the json object that most closely matches the query param
    for i in range(len(data)):
        if similar(data[i][choice_key], request.args['q']) > max:
            max = similar(data[i][choice_key], request.args['q'])
            key = i

    return make_response(jsonify(data[key]), 200) # returning json of the found object

# this is what allows you to run the app
if __name__ == "__main__":
    app.run(debug=True)