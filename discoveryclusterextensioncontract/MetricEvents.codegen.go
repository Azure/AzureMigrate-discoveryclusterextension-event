﻿
// <copyright file="MetricEvents.codegen.go" company="Microsoft Corporation">
// Copyright (c) Microsoft Corporation. All rights reserved.
// </copyright>
// <autogenerated />

package discoveryclusterextensioncontract

type MetricEventId string

    const (
        ApiEncounteredError MetricEventId = "ApiEncounteredError"
        RoleStart MetricEventId = "RoleStart"
        RoleStop MetricEventId = "RoleStop"
        IISApplicationStart MetricEventId = "IISApplicationStart"
        IISApplicationStop MetricEventId = "IISApplicationStop"
    )

//Annotation class
    type ApiEncounteredErrorType struct {
            AnnClientRequestId string
        }

//Method to get AnnotationType
    func ApiEncounteredErrorAnnotation(annClientRequestId string) ApiEncounteredErrorType {
        return ApiEncounteredErrorType{AnnClientRequestId: annClientRequestId}
    }

// Event methods
    func GetApiEncounteredErrorEvent(subscriptionId string, clientRequestId string, errorCode string, annotation ApiEncounteredErrorType) MetricEvent {
        dimensionsDataCollection := map[string]string{}
        dimensionsDataCollection["SubscriptionId"] = subscriptionId
        dimensionsDataCollection["ClientRequestId"] = clientRequestId
        dimensionsDataCollection["ErrorCode"] = errorCode
        
        annotationsDataCollection := map[string]string{}
        annotationsDataCollection["AnnClientRequestId"] = annotation.AnnClientRequestId
 
        return getMetricEvent(ApiEncounteredError, dimensionsDataCollection, annotationsDataCollection)
    }
    func GetRoleStartEvent() MetricEvent {
        dimensionsDataCollection := map[string]string{}
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(RoleStart, dimensionsDataCollection, annotationsDataCollection)
    }
    func GetRoleStopEvent() MetricEvent {
        dimensionsDataCollection := map[string]string{}
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(RoleStop, dimensionsDataCollection, annotationsDataCollection)
    }
    func GetIISApplicationStartEvent() MetricEvent {
        dimensionsDataCollection := map[string]string{}
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(IISApplicationStart, dimensionsDataCollection, annotationsDataCollection)
    }
    func GetIISApplicationStopEvent(hostingEnvShutdownReason string) MetricEvent {
        dimensionsDataCollection := map[string]string{}
        dimensionsDataCollection["HostingEnvShutdownReason"] = hostingEnvShutdownReason
        
        annotationsDataCollection := map[string]string{}
 
        return getMetricEvent(IISApplicationStop, dimensionsDataCollection, annotationsDataCollection)
    }