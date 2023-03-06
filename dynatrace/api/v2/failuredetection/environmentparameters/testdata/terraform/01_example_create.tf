data "dynatrace_request_attribute" "myattr" {
    name = ""
}

resource "dynatrace_failure_detection_parameter" "stupid_404" {
    name = "#name"
    description = "Stupid rule avoiding detecting 404 errors as failure"
    httpResponseCodes {
        server_side_errors = "500-599"
        fail_on_missing_response_code_server_side = true
        client_side_errors = "400-403, 405-498"
        fail_on_missing_response_code_client_side = true
    }
    brokenLinks {
        http404NotFoundFailures = true
        brokenLinkDomains = [
            "example123.com"
        ]
    }
    exception_rules {
        ignore_all_exceptions = true
        custom_error_rules = [
            {
                request_attribute = data.dynatrace_request_attribute.myattr.id,
                condition {
                    compare_operation_type = "NOT_CONTAINS"
                    text_value = "foo"
                    case_sensitive = true
                }
            }
        ]
        ignore_span_failure_detection = true
    }
}