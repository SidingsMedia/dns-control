// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package controller

// type <Service>Controller interface {
// 	<Endpoint>(ctx *gin.Context)
// }

// type <service>Controller struct {
// 	service service.<Service>Service
// }

// func (controller <service>Controller) <Endpoint>(ctx *gin.Context) {
// 	request := &model.<Model>{}
// 	if err := ctx.ShouldBind(request); err != nil && errors.As(err, &validator.ValidationErrors{}){
// 		util.SendBadRequestFieldNames(ctx, err.(validator.ValidationErrors))
// 		return
// 	} else if err != nil {
//         ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GeneralError{
//             Code: http.StatusBadRequest,
//             Message: "Request was malformed",
//         })
//         return
//     }

// 	if err := controller.service.<Handler>(request); err != nil {
//         log.Println(err)
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GeneralError{
//             Code: http.StatusInternalServerError,
//             Message: "An unexpected error occured",
//         })
// 	}
// }

// func New<Service>Controller(engine *gin.Engine, <service>Service service.<Service>Service) {
//     controller := &<service>Controller{
//         service: <service>Service,
//     }
//     api := engine.Group("<service>")
//     {
//         // API Routes go here
//     }
// }
