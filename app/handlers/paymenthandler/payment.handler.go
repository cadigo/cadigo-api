package paymenthandler

import (
	"cadigo-api/app/interface/paymentinterface"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	serv paymentinterface.PaymentService
}

func NewHandler(serv paymentinterface.PaymentService) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (r *Handler) ChillPayCallBack(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	logrus.Info("Params Request : %v", request.URL.Query())
	logrus.Info("Body Request : %v", request.Body)
	logrus.Info("Method Request : %v", request.Method)

	// c.JSON(http.StatusOK, gin.H{"status": "ok", "messsage": "success"})

	// var human Bio
	// err := json.NewDecoder(request.Body).Decode(&human)
	// if err != nil {
	// 	log.Fatalln("There was an error decoding the request body into the struct")
	// }
	// BioData = append(BioData, human)
	// err = json.NewEncoder(writer).Encode(&human)
	// if err != nil {
	// 	log.Fatalln("There was an error encoding the initialized struct")
	// }
}

// // Payment is the resolver for the payment field.
// func (r *Handler) Payment(ctx context.Context, input modelgraph.PaymentInput) (*modelgraph.Payment, error) {
// 	record := modela.Payment{}
// 	err := copier.CopyWithOption(&record, &input, copier.Option{IgnoreEmpty: true})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Insert
// 	if input.ID == nil {
// 		res, err := r.serv.Create(ctx, &record)
// 		if err != nil {
// 			return nil, err
// 		}
// 		c := res.ToGraph()

// 		return &c, nil
// 	} else {
// 		// Update
// 		res, err := r.serv.Update(ctx, *input.ID, &record)
// 		if err != nil {
// 			return nil, err
// 		}
// 		c := res.ToGraph()

// 		return &c, nil
// 	}
// }

// // GetPayment is the resolver for the getPayment field.
// func (r *Handler) GetPayment(ctx context.Context, input modelgraph.GetPaymentInput) (*modelgraph.Payment, error) {
// 	if input.ID != nil {
// 		d, err := r.serv.GetByID(ctx, *input.ID)
// 		if err != nil {
// 			return nil, nil
// 		}

// 		g := d.ToGraph()

// 		return &g, nil
// 	}

// 	return nil, nil
// }
