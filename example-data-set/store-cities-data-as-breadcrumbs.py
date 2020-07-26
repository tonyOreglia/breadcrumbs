import requests
import csv
import json

name_index = 0
lat_index = 1
long_index = 2

def create_notes(notes):
    url = 'http://localhost:80/notes'
    headers = {
        'Content-Type': 'application/json'
    }

    noteChunks = chunks(notes, 1000)


    for chunk in noteChunks:
        payload = json.dumps(chunk)
        print('Saving ' + str(len(chunk)) + ' notes...')
        response = requests.request('POST', url, headers=headers, data = payload)
        print(response.status_code)

def generate_message(row):
    if row[name_index] != '':
        return row[name_index]
    return 'unknown city name'


def lat(row): 
    return location_data(row, lat_index)

def lon(row):
    return location_data(row, long_index)

def location_data(row, index):
    if row[index] == '':
        return ''
    return float(row[index])

def chunks(list, chunk_size):
    for i in range(0, len(list), chunk_size):
        yield list[i:i + chunk_size]

notes = []
with open('cities.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    line_count = 0
    for row in csv_reader:
        if line_count == 0:
            print(f'Column names are {", ".join(row)}')
            line_count += 1
        else:
            if lat(row) != '' and lon(row) != '':
                notes.append({
                    'message': generate_message(row),
                    'latitude': lat(row),
                    'longitude': lon(row),
                    'altitude': -1,
                    'date_created_unix': 1595770552
                })
            line_count += 1
    print(f'Processed {line_count} lines.')

create_notes(notes)
