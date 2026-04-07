import { useState } from "react";
import { CalculatorClient } from "./CalculatorServiceClientPb";
import { CalcRequest } from "./calculator_pb";

const client = new CalculatorClient("http://localhost:8080");
export const App = () => {
  const [num1, setNum1] = useState();
  const [num2, setNum2] = useState();
  const [result, setResult] = useState<number | string>("");
  const handleCalculate = (operation: "add" | "multiple") => {
    const request = new CalcRequest();
    request.setA(num1);
    request.setB(num2);
    if (operation === "add") {
      client.add(request, {}, (err, response) => {
        if (err) {
          console.error("Error performing the addition", err.message);
          setResult("Error:" + err.message);
        }
        setResult(response.getResult());
      });
    } else {
      client.multiply(request,{},(err,response)=>{
        
      })
    }
  };
  return <></>;
};
