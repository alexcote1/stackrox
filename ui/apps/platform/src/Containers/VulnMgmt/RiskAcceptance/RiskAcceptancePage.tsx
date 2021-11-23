import React, { ReactElement } from 'react';
import {
    Divider,
    Flex,
    FlexItem,
    PageSection,
    PageSectionVariants,
    Tab,
    TabContent,
    Tabs,
    TabTitleText,
    TextContent,
    Title,
} from '@patternfly/react-core';

import useTabs from 'hooks/patternfly/useTabs';

import PendingApprovals from './PendingApprovals';
import ApprovedDeferrals from './ApprovedDeferrals';
import ApprovedFalsePositives from './ApprovedFalsePositives';

const TABS = {
    PENDING_APPROVALS: 'Pending Approvals',
    APPROVED_DEFERRALS: 'Approved Deferrals',
    APPROVED_FALSE_POSITIVES: 'Approved False Positives',
};

function RiskAcceptancePage(): ReactElement {
    const { activeKeyTab, onSelectTab } = useTabs({ defaultTab: TABS.PENDING_APPROVALS });

    return (
        <>
            <PageSection variant={PageSectionVariants.light}>
                <Flex
                    alignItems={{
                        default: 'alignItemsFlexStart',
                        md: 'alignItemsCenter',
                    }}
                    direction={{ default: 'column', md: 'row' }}
                    flexWrap={{ default: 'nowrap' }}
                    spaceItems={{ default: 'spaceItemsXl' }}
                >
                    <FlexItem grow={{ default: 'grow' }}>
                        <TextContent>
                            <Title headingLevel="h1">Risk Acceptance</Title>
                        </TextContent>
                    </FlexItem>
                </Flex>
            </PageSection>
            <Divider component="div" />
            <Tabs activeKey={activeKeyTab} onSelect={onSelectTab}>
                <Tab
                    eventKey={TABS.PENDING_APPROVALS}
                    tabContentId={TABS.PENDING_APPROVALS}
                    title={<TabTitleText>{TABS.PENDING_APPROVALS}</TabTitleText>}
                />
                <Tab
                    eventKey={TABS.APPROVED_DEFERRALS}
                    tabContentId={TABS.APPROVED_DEFERRALS}
                    title={<TabTitleText>{TABS.APPROVED_DEFERRALS}</TabTitleText>}
                />
                <Tab
                    eventKey={TABS.APPROVED_FALSE_POSITIVES}
                    tabContentId={TABS.APPROVED_FALSE_POSITIVES}
                    title={<TabTitleText>{TABS.APPROVED_FALSE_POSITIVES}</TabTitleText>}
                />
            </Tabs>
            <TabContent
                eventKey={TABS.PENDING_APPROVALS}
                id={TABS.PENDING_APPROVALS}
                hidden={activeKeyTab !== TABS.PENDING_APPROVALS}
            >
                <PendingApprovals />
            </TabContent>
            <TabContent
                eventKey={TABS.APPROVED_DEFERRALS}
                id={TABS.APPROVED_DEFERRALS}
                hidden={activeKeyTab !== TABS.APPROVED_DEFERRALS}
            >
                <ApprovedDeferrals />
            </TabContent>
            <TabContent
                eventKey={TABS.APPROVED_FALSE_POSITIVES}
                id={TABS.APPROVED_FALSE_POSITIVES}
                hidden={activeKeyTab !== TABS.APPROVED_FALSE_POSITIVES}
            >
                <ApprovedFalsePositives />
            </TabContent>
        </>
    );
}

export default RiskAcceptancePage;
