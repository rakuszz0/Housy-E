import React, { useContext, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import NavbarProject from "../components/NavbarProject";
import Icon from "../components/Icon.svg";
import Row from "react-bootstrap/esm/Row";
import Col from "react-bootstrap/esm/Col";
import Button from "react-bootstrap/esm/Button";
import Container from "react-bootstrap/esm/Container";
import Elipse from "../assets/img/Ellipse 7.png";
import Elipsee from "../assets/img/Ellipse 8.png";
import Line from "../assets/img/Line 9.png";
import CodeQr from "../assets/img/woyy.png";
import { useNavigate, useParams } from "react-router-dom";
import { UserContext } from "../context/userContext";
import { useMutation, useQuery } from "react-query";
import { API } from "../config/api";
import Moment from "react-moment";
import { convert } from "rupiah-format";
import NavbarWithoutSearch from "../components/NavbarWithoutSearch";

export default function MyBooking(props) {
  const getData = JSON.parse(localStorage.getItem("check_in"));

  // const getToken = localStorage.getItem("token");
  let history = useNavigate();
  // const hasilDecode = jwt(getToken);



  const { id } = useParams();

  const [state, dispatch] = useContext(UserContext);

  // fetching data house from database
  let { data: house, refetch } = useQuery("detailCache", async () => {
    const config = {
      method: "GET",
      headers: {
        Authorization: "Basic " + localStorage.token,
      },
    };
    const response = await API.get("/house/" + id, config);
    console.log("data response test", response);
    return response.data.data;
  });

  useEffect(() => {
    //change this to the script source you want to load, for example this is snap.js sandbox env
    const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
    //change this according to your client-key
    const myMidtransClientKey = process.env.REACT_APP_MIDTRANS_CLIENT_KEY;

    let scriptTag = document.createElement("script");
    scriptTag.src = midtransScriptUrl;
    // optional if you want to set script attribute
    // for example snap.js have data-client-key attribute
    scriptTag.setAttribute("data-client-key", myMidtransClientKey);

    document.body.appendChild(scriptTag);
    return () => {
      document.body.removeChild(scriptTag);
    };
  }, []);

  const dateTime = new Date();
  const checkin = new Date(getData.check_in);
  const checkout = new Date(getData.check_out);

  console.log(house.type_rent, "type")
  let p = 0
  if (house.type_rent === "Day") { //Day
    p = house?.price * Math.floor((checkout - checkin) / (1000 * 60 * 60 * 24))
  } else if (house.type_rent === "Month") { //Month
    p = house?.price * Math.floor((checkout - checkin) / (1000 * 60 * 60 * 24 * 30))
  } else if (house.type_rent === "Year") { //Year
    p = house?.price * Math.floor((checkout - checkin) / (1000 * 60 * 60 * 24 * 30 * 12))
  }

  const handleTransaction = useMutation(async () => {
    try {
      const response = await API.post("/transaction", {
        check_in: checkin,
        check_out: checkout,
        house_id: house.id,
        user_id: state.user.id,
        total: p,
        status_payment: "Pending",
        attachment: "image.png",
        created_at: dateTime,
        updated_at: dateTime,
      });

      const tokenBaru = response.data.data.token;
      console.log("habis add transaction tokennnnnn : ", response);

      // const token = response.data.data.token;
      console.log("ini tokennnnn", response);
      console.log("ini tokennnnnbaru", tokenBaru);

      window.snap.pay(tokenBaru, {
        onSuccess: function (result) {
          /* You may add your own implementation here */
          console.log(result);
          history("/history");
        },
        onPending: function (result) {
          /* You may add your own implementation here */
          console.log(result);
          history("/my-booking");
        },
        onError: function (result) {
          /* You may add your own implementation here */
          console.log(result);
        },
        onClose: function () {
          /* You may add your own implementation here */
          alert("you closed the popup without finishing the payment");
        },
      });
    } catch (error) {
      console.log(error);
    }
  });

  useEffect(() => {
    //change this to the script source you want to load, for example this is snap.js sandbox env
    const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
    //change this according to your client-key
    const myMidtransClientKey = "SB-Mid-client-PXZXQGaKnNSLWukm";

    let scriptTag = document.createElement("script");
    scriptTag.src = midtransScriptUrl;
    // optional if you want to set script attribute
    // for example snap.js have data-client-key attribute
    scriptTag.setAttribute("data-client-key", myMidtransClientKey);

    document.body.appendChild(scriptTag);
    return () => {
      document.body.removeChild(scriptTag);
    };
  }, []);

  return (
    <>
      <NavbarWithoutSearch />
      <Container className="myc fmb" style={{ width: "60%", marginTop: "200px", marginBottom: "50px" }}>
        <div className="border border-3 p-4 pe-0 pb-0">
          <Row style={{}} className="d-flex jcb">
            <Col className="" md="auto" lg={4}>
              <img src={Icon} alt="" />
            </Col>
            <Col className="" md="auto" lg={4}>
              <h2 className="text-center p-0 m-0 fw-bold">Booking</h2>
              <p className="text-center p-0 m-0">
                <Moment format="dddd" className="fw-bold">
                  {dateTime}
                </Moment>
                , <Moment format="D MMM YYYY">{dateTime}</Moment>
              </p>
            </Col>
          </Row>
          <Row style={{}} className="d-flex jcb align-items-center pb-3">
            <Col className="" md="auto" lg={4}>
              <h5 className="fw-bold">{house?.name}</h5>
              <p>{house?.address}</p>
              <p className="bg-danger w-50 text-center p-1 bg-opacity-10 text-danger">Waiting Payment</p>
            </Col>
            <Col className="" md="auto" lg={4}>
              <div className="d-flex flex-column ">
                <div className="d-flex  align-items-center gap-4">
                  <div>
                    <img src={Elipse} alt="" />
                  </div>
                  <div className="d-flex flex-column">
                    <span>Check-in</span>
                    <span>
                      <Moment format="DD MMM YYYY">{getData.check_in}</Moment>
                    </span>
                  </div>
                  <div className="ms-3 d-flex flex-column">
                    <span>Amenities</span>
                    <span>{house?.amenities}</span>
                  </div>
                </div>

                <div className="d-flex ">
                  <img style={{ marginLeft: "6px" }} src={Line} alt="" />
                </div>
                <div className="d-flex  align-items-center gap-4">
                  <div>
                    <img src={Elipsee} alt="" />
                  </div>

                  <div className="d-flex flex-column ">
                    <span>Check-Out</span>
                    <span>
                      <Moment format="DD MMM YYYY">{getData.check_out}</Moment>
                    </span>
                  </div>
                  <div className="ms-3 d-flex flex-column ">
                    <span>Type of Rent</span>
                    <span>{house?.type_rent}</span>
                  </div>
                </div>
              </div>
            </Col>
            <Col className="d-flex flex-column justify-content-center align-items-center gap-2" md="auto" lg={4}>
              <img src={CodeQr} alt="" style={{ width: 150 }} />
            </Col>
          </Row>
          <Row className="d-flex">
            <Row>
              <Col className="d-flex" md="auto" lg={8}>
                <Col className="d-flex align-items-center" md="auto" lg={1}>
                  <p className="m-0 py-2">No</p>
                </Col>
                <Col className="d-flex align-items-center" md="auto" lg={3}>
                  <p className="m-0">Full Name</p>
                </Col>
                <Col className="d-flex align-items-center" md="auto" lg={3}>
                  <p className="m-0">Gender</p>
                </Col>
                <Col className="d-flex align-items-center" md="auto" lg={3}>
                  <p className="m-0">Phone</p>
                </Col>
              </Col>
            </Row>
            <Row className="border border-start-0 border-end-0  ">
              <Col className="d-flex" lg={8}>
                <Col className="d-flex align-items-center" md="auto" lg={1}>
                  <p className="m-0">1</p>
                </Col>
                <Col className="d-flex align-items-center" md="auto" lg={3}>
                  <p className="m-0">{state.user.fullname}</p>
                </Col>
                <Col className="d-flex align-items-center" md="auto" lg={3}>
                  <p className="m-0">{state.user.gender}</p>
                </Col>
                <Col className="d-flex align-items-center" md="auto" lg={3}>
                  <p className="m-0">{state.user.phone}</p>
                </Col>
              </Col>
              <Col className="d-flex align-items-center">
                <p className="ps-3 m-0">Long time rent</p>
              </Col>
              <Col className="d-flex align-items-center">
                <p className="m-0 py-2">
                  : <Moment duration={getData.check_in} date={getData.check_out} />
                </p>
              </Col>
            </Row>
            <Row className="justify-content-end">
              <Col className="d-flex align-items-center" lg={2}>
                <p className=" m-0 ps-3 py-2">Total</p>
              </Col>
              <Col className="d-flex align-items-center" lg={2}>
                <p className="m-0 text-danger fw-bold">: {p.toLocaleString("id-ID", { style: "currency", currency: "IDR" })}</p>
              </Col>
            </Row>
          </Row>
        </div>
        <div className="d-flex justify-content-end">
          <Button type="submit" style={{ width: "200px" }} onClick={() => handleTransaction.mutate()}>
            Pay
          </Button>
        </div>
      </Container>
    </>
  );
}
