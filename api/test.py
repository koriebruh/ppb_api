import requests

# URL API yang ingin diakses
url = "https://api-ppb.vercel.app/api/products"

try:
    # Mengirim permintaan GET ke API
    response = requests.get(url)
    status_code = response.status_code

    # Menampilkan status kode dan data jika berhasil
    if status_code == 200:
        data = response.json()
        result = {"status_code": status_code, "data": data}
    else:
        result = {"status_code": status_code, "error": response.text}
except Exception as e:
    # Menampilkan error jika terjadi kegagalan
    result = {"error": str(e)}

result
