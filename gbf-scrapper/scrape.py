from bs4 import BeautifulSoup # HTML Parser
import psycopg2 # Postgres DB
import requests # cURL
from dotenv import load_dotenv # Environment Library

base_wiki_url = 'https://gbf.wiki/Weapon_Lists/SSR/Fire' # Omit the Fire when I get to actual
game_grab = 'game.granbluefantasy.jp/#list'

load_dotenv()

def grab_single_from_game():
    source = base_wiki_url
    soup = BeautifulSoup(source, 'lxml')

    # Note that everything I need is found in the id='cnt-detail' class
    # For skills, if there is multiple of them, they exist in their own prt box.

    weapon_name = soup.find_all(class_='txt-item-name')
    weapon_element_and_type = soup.find_all(class_='prt-attr-weapon') # Note the first is element, then weapon type
    weapon_base_stat_hp = soup.find_all(class_='txt-hp-value')
    weapon_base_stat_atk = soup.find_all(class_='txt-atk-value')

    weapon_ougi = soup.find_all('prt-detail-special')
    weapon_skills = soup.find_all('prt-detail-skill')

def grab_single_from_wiki(url, element):
    source = requests.get(url)
    soup = BeautifulSoup(source.text, 'html.parser')

    # Grab table
    wep_table = soup.find_all('table')[1]

    # Store into a list
    wep_table_body = wep_table.find_all('tbody')[0]
    for row in wep_table_body.find_all('tr'):
        print('====================')
        columns = row.find_all('td')
        if columns:
            columns = [ele.text.strip() for ele in columns]
            columns[8] = spaces_after_sentences(columns[8])
            columns[2] = 'SSR'
            columns[3] = element
            print(columns)
            insert_single_to_database(columns)
        
        # First td element has picture
        # Second td element has the name
        # Third td is the rarity (fixed)
        # Fourth td is element (fixed)
        # Fifth td is weapon type
        # Six td is attack
        # Seventh is hp
        # Eigth is weapon skill pictures
        # Ninth has weapon skill info and ougi info

    #wep_name = soup.find_all(class_='char-name')[0].get_text()
    

def insert_batch_to_database(wep_list):
    try:
        conn = psycopg2.connect("")
        curr = conn.cursor()

        for wep in wep_list:
            curr.execute("INSERT INTO weapons VALUES (%s, %s, %s)", (wep.name, wep.type, 'fire'))

        conn.commit()
        conn.close()
    except:
        print('Could not batch insert to database')

def insert_single_to_database(wep):
    try:
        conn = psycopg2.connect("")
        curr = conn.cursor()
        try:
            # TODO: Change this when i settle on the schema
            curr.execute("""INSERT INTO weapons VALUES (DEFAULT, %s, 'katana', %s)""", (wep[1], wep[3]))
        except Exception as ex:
            print('Could not insert data into the weapons table')
            print(ex)

        # =================================================
        # TESTING if item was put into database properly
        try:
            curr.execute("""SELECT * FROM weapons""")
            rows = curr.fetchall()
            print('Rows:')
            for row in rows:
                print(row)
        except Exception as ex:
            print('Cannot retrieve info from database.')
            print(ex)

        # =================================================
        # uncomment this when i want the database to save properly
        #conn.commit()
        conn.close()
    except:
        print('I am unable to connect to the database.')

# Helper Function
def print_database(db_ptr):
    try:
        db_ptr.execute("""SELECT * FROM weapons""")

        rows = db_ptr.fetchall()
        print('Rows:')
        for row in rows:
            print(row)
    except:
        print('Something went wrong with trying to access the db')

# Sanitize the string. Ensuring that sentence spacing is enforced.
def spaces_after_sentences(s):
    for i in range(len(s) - 1):
        if s[i] == '.' and not s[i+1].isdigit() and s[i+1] != ' ':
            s = s[:i] + '\n' + s[i+1:]

    return s.replace('\n', '. ')


# Main 
scrape_details = [('https://gbf.wiki/Weapon_Lists/SSR/Fire', 'fire'), ('https://gbf.wiki/Weapon_Lists/SSR/Water', 'water'), ('https://gbf.wiki/Weapon_Lists/SSR/Earth', 'earth'), ('https://gbf.wiki/Weapon_Lists/SSR/Wind', 'wind'), ('https://gbf.wiki/Weapon_Lists/SSR/Light', 'light'), ('https://gbf.wiki/Weapon_Lists/SSR/Dark', 'dark')]

grab_single_from_wiki(scrape_details[0][0], scrape_details[0][1])
#for (url, ele) in scrape_details:
#    grab_single_from_wiki(url, ele)
