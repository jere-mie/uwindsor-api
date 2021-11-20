from flask import Flask, request, jsonify, make_response, render_template, url_for, flash, redirect
import json
from difflib import SequenceMatcher
from flask_cors import CORS, cross_origin
import sys
import os

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

# home webpage
@app.route('/home')
@app.route('/')
def home():
    return render_template('home.html')

# guide page
@app.route('/guide')
def guide():
    return render_template('guide.html')

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

    jsondata = "" # getting json from file
    with open(f'static/datasets/{choice}/{choice}.json', 'r') as f:
        jsondata = json.loads(f.read())

    if 'q' not in request.args: # we return all data if no param
        return make_response(jsonify(jsondata), 200)

    choice_key = keys[choice] # getting the name of the attribute to query by in the json
    key = 0 # the key of the data found
    max = -1 # max similarity between query and each json item

    # finding the json object that most closely matches the query param
    for i in range(len(jsondata)):
        if similar(jsondata[i][choice_key], request.args['q']) > max:
            max = similar(jsondata[i][choice_key], request.args['q'])
            key = i

    return make_response(jsonify(jsondata[key]), 200) # returning json of the found object


# running site
if __name__=='__main__':
    # run this command with any additional arg to run in production
    if len(sys.argv) > 1:
        print('<< PROD >>')
        os.system(f"gunicorn -b '127.0.0.1:{data['port']}' website:app")
    # or just run without an additional arg to run in debug
    else:
        print('<< DEBUG >>')
        app.run(debug=True)
