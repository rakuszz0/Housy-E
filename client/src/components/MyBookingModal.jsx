import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useNavigate, useParams } from "react-router-dom";
import { useState } from "react";
import { useMutation } from "react-query";
import Moment from "react-moment";

export default function MyBookingModal(props) {
  const { id } = useParams();
  const navigate = useNavigate();
  // const month = props.detail.type_rent
  // const year = props.detail.type_rent

  const [checkIn, setCheckIn] = useState({
    check_in: "",
    check_out: "",
  });

  const handleOnChange = (e) => {
    setCheckIn({
      ...checkIn,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();
      // const check_inDate = new Date(checkIn.check_in)
      // const check_outDate = new Date(checkIn.check_out)
      // const duration = (check_outDate - check_inDate) / (1000 * 60 * 60)
      const formData = new FormData();
      formData.append("check_in", checkIn.check_in);
      formData.append("check_out", checkIn.check_out);

      localStorage.setItem("check_in", JSON.stringify(checkIn));
      navigate(`/my-booking/${id}`);

      <alert variant="danger" className="py-1">
        succses
      </alert>;

      // Handling response here
    } catch (error) {
      <alert variant="danger" className="py-1">
        Failed
      </alert>;
      console.log(error);
    }
  });

  console.log("apaaa", props.month);
  //   <Form.Group className="mb-3" style={{ display: "flex", flexDirection: "column" }} controlId="exampleForm.ControlInput1">

  //   <Form.Label style={{ fontWeight: "bold" }}>Check-in</Form.Label>
  //   <Form.Control type="date" name="check_in" value={checkIn.check_in} onChange={handleOnChange} />
  // </Form.Group>

  // <Form.Group className="mb-3" style={{ display: "flex", flexDirection: "column" }} controlId="exampleForm.ControlInput1">
  //   <Form.Label style={{ fontWeight: "bold" }}>Check-Out</Form.Label>
  //   <Form.Control type="date" name="check_out" value={checkIn.check_out} onChange={handleOnChange} />
  // </Form.Group>

  return (
    <Modal {...props} aria-labelledby="contained-modal-title-vcenter" centered>
      <h3 className="fw-bold text-center my-5">How long you will stay</h3>
      <Modal.Body>
        <Form>
          {props.month === "Day" ? (<> <Form.Group className="mb-3" style={{ display: "flex", flexDirection: "column" }} controlId="exampleForm.ControlInput1">

            <Form.Label style={{ fontWeight: "bold" }}>Check-in</Form.Label>
            <Form.Control type="date" name="check_in" value={checkIn.check_in} onChange={handleOnChange} />
          </Form.Group>

            <Form.Group className="mb-3" style={{ display: "flex", flexDirection: "column" }} controlId="exampleForm.ControlInput1">
              <Form.Label style={{ fontWeight: "bold" }}>Check-Out</Form.Label>
              <Form.Control type="date" name="check_out" value={checkIn.check_out} onChange={handleOnChange} />
            </Form.Group></>) : (<> <Form.Group className="mb-3" style={{ display: "flex", flexDirection: "column" }} controlId="exampleForm.ControlInput1">

              <Form.Label style={{ fontWeight: "bold" }}>Check-in</Form.Label>
              <Form.Control type="month" name="check_in" value={checkIn.check_in} onChange={handleOnChange} />
            </Form.Group>

              <Form.Group className="mb-3" style={{ display: "flex", flexDirection: "column" }} controlId="exampleForm.ControlInput1">
                <Form.Label style={{ fontWeight: "bold" }}>Check-Out</Form.Label>
                <Form.Control type="month" name="check_out" value={checkIn.check_out} onChange={handleOnChange} />
              </Form.Group></>)}
          <Button
            onClick={(e) => {
              handleSubmit.mutate(e);
            }}
          >
            Order
          </Button>
        </Form>
      </Modal.Body>
    </Modal>
  );
}
