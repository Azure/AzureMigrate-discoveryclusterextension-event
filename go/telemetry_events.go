﻿
// Code generated by controller-gen. DO NOT EDIT.
// <copyright file="TelemetryEvent.go" company="Microsoft Corporation">
// Copyright (c) Microsoft Corporation. All rights reserved.
// </copyright>
// <autogenerated />

package discoveryextensionevents

    type TelemetryEventId int

    const (
        EntityDiscovered TelemetryEventId = 2002
)

        func EntityDiscoveredTelemetryEvent(entityId string, entityType string, entityProperties string, executionTimeSec string, entityErrors string) TelemetryEvent {
            return getTelemetryEvent (
              EntityDiscovered,
              map[string]string {

                    "EntityId": entityId,
                    "EntityType": entityType,
                    "EntityProperties": entityProperties,
                    "ExecutionTimeSec": executionTimeSec,
                    "EntityErrors": entityErrors,
            })
        }
