import { View } from 'react-native'
import React from 'react'

export class Picker extends React.Component {
  constructor (props) {
    super(props)

    this.props.navigation.navigate(
      // TODO: when will find a way to use all our components in Library, implement skip of onboarding in test replace the next line by:
      // this.props.screenProps.firstLaunch || process.env['ENVIRONMENT'] === 'integration_test'
      this.props.navigation.getParam('firstLaunch', false) && process.env['ENVIRONMENT'] !== 'integration_test'
        ? 'switch/onboarding'
        : 'switch/main'
    )
  }

  render () {
    return <View />
  }
}