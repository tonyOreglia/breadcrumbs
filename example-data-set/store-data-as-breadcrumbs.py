import requests
import csv
import json


def create_notes(notes):
    url = 'http://localhost:80/notes'
    headers = {
    'Content-Type': 'application/json'
    }

    noteChunks = chunks(notes, 1000)


    for chunk in noteChunks:
        payload = json.dumps(chunk)
        print('Saving ' + str(len(chunk)) + 'notes...')
        response = requests.request('POST', url, headers=headers, data = payload)
        print(response.status_code)

def generate_message(row):
    if row[1] != '':
        return row[1] + ', ' + row[0] + ': ' + death_count(row)
    return row[0] + ': ' + death_count(row)


def death_count(row):
    if row[-2] == '':
        return '0'
    return str(int(float(row[-2])))

def lat(row): 
    return location_data(row, 2)

def lon(row):
    return location_data(row, 3)

def location_data(row, index):
    if row[index] == '':
        return ''
    return float(row[index])

def chunks(list, chunk_size):
    for i in range(0, len(list), chunk_size):
        yield list[i:i + chunk_size]

notes = []
with open('covid-19-all.csv') as csv_file:
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
                    'longitude': lon(row)
                })
            line_count += 1
    print(f'Processed {line_count} lines.')

create_notes(notes)
