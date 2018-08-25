// eslint-disable-next-line no-unused-vars
import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import PropTypes from 'prop-types';
import { get, pick } from 'lodash';
import { reduxForm, FormSection, getFormValues } from 'redux-form';

import { PanelSwaggerField, editablePanelify } from 'shared/EditablePanel';
import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';

import './index.css';

const surveyFields = [
  'pm_survey_pack_date',
  'pm_survey_pickup_date',
  'pm_survey_latest_pickup_date',
  'pm_survey_earliest_delivery_date',
  'pm_survey_latest_delivery_date',
  'pm_survey_weight_estimate',
  'pm_survey_progear_weight_estimate',
  'pm_survey_spouse_progear_weight_estimate',
  'pm_survey_notes',
];

const SurveyDisplay = props => {
  const fieldProps = {
    schema: props.shipmentSchema,
    values: props.shipment,
  };

  return (
    <React.Fragment>
      <div className="editable-panel-3-column">
        <PanelSwaggerField
          title="Pack Date"
          fieldName="pm_survey_pack_date"
          nullWarning
          {...fieldProps}
        />
        <PanelSwaggerField
          title="Pickup Date"
          fieldName="pm_survey_pickup_date"
          nullWarning
          {...fieldProps}
        />
        <PanelSwaggerField
          title="Latest Pickup Date"
          fieldName="pm_survey_latest_pickup_date"
          nullWarning
          {...fieldProps}
        />
      </div>
      <div className="editable-panel-3-column">
        <PanelSwaggerField
          title="Earliest Pickup Date"
          fieldName="pm_survey_earliest_delivery_date"
          nullWarning
          {...fieldProps}
        />
        <PanelSwaggerField
          title="Latest Delivery Date"
          fieldName="pm_survey_latest_delivery_date"
          nullWarning
          {...fieldProps}
        />
      </div>
      <div className="editable-panel-3-column">
        <PanelSwaggerField
          title="Weight Estimate"
          fieldName="pm_survey_weight_estimate"
          nullWarning
          {...fieldProps}
        />
        <PanelSwaggerField
          title="Progear Weight Estimate"
          fieldName="pm_survey_progear_weight_estimate"
          {...fieldProps}
        />
        <PanelSwaggerField
          title="Spouse Progear Weight Estimate"
          fieldName="pm_survey_spouse_progear_weight_estimate"
          {...fieldProps}
        />
      </div>
      <div className="editable-panel-column">
        <PanelSwaggerField
          title="Notes"
          fieldName="pm_survey_notes"
          {...fieldProps}
        />
      </div>
    </React.Fragment>
  );
};

const SurveyEdit = props => {
  const schema = props.shipmentSchema;
  return (
    <React.Fragment>
      <FormSection name="survey">
        <div className="editable-panel-3-column">
          <SwaggerField
            fieldName="pm_survey_pack_date"
            swagger={schema}
            className="half-width"
            required
          />
          <SwaggerField
            fieldName="pm_survey_pickup_date"
            swagger={schema}
            className="half-width"
            required
          />
          <SwaggerField
            fieldName="pm_survey_latest_pickup_date"
            swagger={schema}
            className="half-width"
            required
          />
        </div>

        <div className="editable-panel-3-column">
          <SwaggerField
            fieldName="pm_survey_earliest_delivery_date"
            swagger={schema}
            className="half-width"
            required
          />
          <SwaggerField
            fieldName="pm_survey_latest_delivery_date"
            swagger={schema}
            className="half-width"
            required
          />
        </div>

        <div className="editable-panel-3-column">
          <SwaggerField
            fieldName="pm_survey_weight_estimate"
            swagger={schema}
            className="half-width"
            required
          />
          <SwaggerField
            fieldName="pm_survey_progear_weight_estimate"
            swagger={schema}
            className="half-width"
          />
          <SwaggerField
            fieldName="pm_survey_spouse_progear_weight_estimate"
            swagger={schema}
            className="half-width"
          />
        </div>
        <SwaggerField
          fieldName="pm_survey_notes"
          swagger={schema}
          className="half-width"
        />
      </FormSection>
    </React.Fragment>
  );
};

const formName = 'shipment_pre_move_survey';

let PremoveSurveyPanel = editablePanelify(SurveyDisplay, SurveyEdit);
PremoveSurveyPanel = reduxForm({
  form: formName,
  enableReinitialize: true,
  keepDirtyOnReinitialize: true,
})(PremoveSurveyPanel);

PremoveSurveyPanel.propTypes = {
  shipment: PropTypes.object,
};

function mapStateToProps(state, props) {
  let formValues = getFormValues(formName)(state);

  return {
    // reduxForm
    formValues: formValues,
    initialValues: {
      survey: pick(props.shipment, surveyFields),
    },

    shipmentSchema: get(state, 'swagger.spec.definitions.Shipment', {}),

    hasError: !!props.error,
    errorMessage: props.error,
    isUpdating: false,

    // editablePanelify
    getUpdateArgs: function() {
      return [get(props, 'shipment.id'), formValues.survey];
    },
  };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators({}, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(PremoveSurveyPanel);
