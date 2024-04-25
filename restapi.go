
package main
import (
 "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    )

    type data struct {
      Endpoint       string                   `json:'endpoint'`
      Quotes         []map[string]interface{} `json:'quotes'`
      Requested_time string               `json:'requested_time'`
      Timestamp      int32                    `json:'timestamp'`
  }

  func main(){
  currencies := "EURUSD,GBPUSD"
  api_key := "N5vDPdrDqz6WDI0poAT0"
  url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key
  resp, getErr := http.Get(url)
  if getErr != nil {
      log.Fatal(getErr)
  }
  defer resp.Body.Close() // Don't forget to close the response body when done with it
  body, readErr := ioutil.ReadAll(resp.Body)
  if readErr != nil {
      log.Fatal(readErr)
  }
  fmt.Println(string(body))

   data_obj := data{}

     jsonErr := json.Unmarshal(body, &data_obj)
        if jsonErr != nil {
           log.Fatal(jsonErr)
        }

      fmt.Println("endpoint", data_obj.Endpoint, "requested time", data_obj.Requested_time, "timestamp", data_obj.Timestamp)

      for key, value := range data_obj.Quotes {
          fmt.Println(key)
          baseCurrency, baseOK := value["base_currency"].(string)
          quoteCurrency, quoteOK := value["quote_currency"].(string)
          bid, bidOK := value["bid"].(float64)
          ask, askOK := value["ask"].(float64)
          mid, midOK := value["mid"].(float64)

          if baseOK && quoteOK && bidOK && askOK && midOK {
              fmt.Println("symbol", baseCurrency+quoteCurrency, "bid", bid, "ask", ask, "mid", mid)
          } else {
              fmt.Println("Error: Invalid data format in quotes map")
          }
      }
}