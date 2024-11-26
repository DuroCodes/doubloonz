import re
import json


def extract_data(html: str) -> str:
    cost_pattern = r'<span class="text-green-500[^>]*><img[^>]*>(\d+)</span>'
    costs = re.findall(cost_pattern, html)

    name_pattern = (
        r'<h3 class="tracking-tight text-xl font-bold text-center">([^<]+)</h3>'
    )
    names = re.findall(name_pattern, html)

    items = [
        {"name": name.strip(), "cost": int(cost)} for name, cost in zip(names, costs)
    ]

    return json.dumps(items)


with open("data.html") as file:
    html = file.read()
    print(extract_data(html))
