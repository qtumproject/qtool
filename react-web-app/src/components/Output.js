import { Card , CardTitle, ListGroup, ListGroupItem} from "reactstrap";

const Output = ({ output }) => {
  if (output === null) {
    return <div className="response">Click submit to see result</div>;
  }
  if (output.error) {
    let error = output.error;
    return <div className="response">
      Error: {error.message !== null ? error.message : error} <br />
      Verify input and try again...
      </div>;
  }
  return (
      <div className="response">
        {Object.entries(output.result).map(([key, value]) => (
           <div key={key}>
            <Card
              style={{
                width: "100%",
                textAlign: "left",
              }}
            >
              <CardTitle 
              style={{
                color: "#f1356d",
              }}
              >
                {key}
              </CardTitle>
              <ListGroup flush>
                <ListGroupItem>
                  {typeof value === "object" ? JSON.stringify(value, null, 2) : value}
                </ListGroupItem>
              </ListGroup>
            </Card>
          </div>
        ))}
      </div>

    );
}

export default Output;