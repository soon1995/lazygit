package i18n

func dutchTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "Niet genoeg ruimte om de panelen te renderen",
		DiffTitle:                           "Diff",
		FilesTitle:                          "Bestanden",
		BranchesTitle:                       "Branches",
		CommitsTitle:                        "Commits",
		StashTitle:                          "Stash",
		UnstagedChanges:                     "Unstaged wijzigingen",
		StagedChanges:                       "Staged wijzigingen",
		MainTitle:                           "Hoofd",
		StagingTitle:                        "Staging",
		NormalTitle:                         "Normaal",
		CommitSummary:                       "Commitbericht",
		CredentialsUsername:                 "Gebruikersnaam",
		CredentialsPassword:                 "Wachtwoord",
		CredentialsPassphrase:               "Voer een wachtwoordzin in voor de SSH-sleutel",
		PassUnameWrong:                      "Wachtwoord en/of gebruikersnaam verkeerd",
		CommitChanges:                       "Commit veranderingen",
		AmendLastCommit:                     "Wijzig laatste commit",
		AmendLastCommitTitle:                "Wijzig laatste commit",
		SureToAmend:                         "Weet je zeker dat je de laatste commit wilt wijzigen? U kunt het commit-bericht wijzigen vanuit het commits-paneel.",
		NoCommitToAmend:                     "Er is geen commits om te wijzigen.",
		CommitChangesWithEditor:             "Commit veranderingen met de git editor",
		StatusTitle:                         "Status",
		Menu:                                "Menu",
		Execute:                             "Uitvoeren",
		ToggleStaged:                        "Toggle staged",
		ToggleStagedAll:                     "Toggle staged alle",
		Refresh:                             "Verversen",
		Push:                                "Push",
		Pull:                                "Pull",
		Scroll:                              "Scroll",
		FilterStagedFiles:                   "Show only staged files",
		FilterUnstagedFiles:                 "Show only unstaged files",
		ResetCommitFilterState:              "Reset commit file state filter",
		MergeConflictsTitle:                 "Merge conflicten",
		Checkout:                            "Uitchecken",
		PullWait:                            "Pullen...",
		PushWait:                            "Pushen...",
		FetchWait:                           "Fetchen...",
		SoftReset:                           "Zacht reset",
		AlreadyCheckedOutBranch:             "Je hebt deze branch al uitgecheckt",
		SureForceCheckout:                   "Weet je zeker dat je het uitchecken wil forceren? Al je lokale verandering zullen worden verwijdert",
		ForceCheckoutBranch:                 "Forceer uitchecken op deze branch",
		BranchName:                          "Branch naam",
		NewBranchNameBranchOff:              "Nieuw branch naam (Branch is afgeleid van '{{.branchName}}')",
		CantDeleteCheckOutBranch:            "Je kan een uitgecheckte branch niet verwijderen!",
		DeleteBranch:                        "Verwijder branch",
		DeleteBranchMessage:                 "Weet je zeker dat je branch '{{.selectedBranchName}}' wilt verwijderen?",
		ForceDeleteBranchMessage:            "Weet je zeker dat je branch '{{.selectedBranchName}}' geforceerd wil verwijderen?",
		RebaseBranch:                        "Rebase branch",
		CantRebaseOntoSelf:                  "Je kan niet een branch rebasen op zichzelf",
		CantMergeBranchIntoItself:           "Je kan niet een branch in zichzelf mergen",
		ForceCheckout:                       "Forceer checkout",
		CheckoutByName:                      "Uitchecken bij naam",
		NewBranch:                           "Nieuwe branch",
		NoBranchesThisRepo:                  "Geen branches voor deze repo",
		CommitWithoutMessageErr:             "Je kan geen commit maken zonder commit bericht",
		CloseCancel:                         "Sluiten",
		Confirm:                             "Bevestig",
		Close:                               "Sluiten",
		Quit:                                "Quit",
		SquashDown:                          "Squash beneden",
		FixupCommit:                         "Fixup commit",
		CannotSquashOrFixupFirstCommit:      "There's no commit below to squash into",
		Fixup:                               "Fixup",
		SureFixupThisCommit:                 "Weet je zeker dat je fixup wil uitvoeren op deze commit? De commit hieronder zol worden squashed in deze",
		SureSquashThisCommit:                "Weet je zeker dat je deze commit wil samenvoegen met de commit hieronder?",
		Squash:                              "Squash",
		PickCommit:                          "Kies commit (wanneer midden in rebase)",
		RevertCommit:                        "Commit ongedaan maken",
		RewordCommit:                        "Hernoem commit",
		DeleteCommit:                        "Verwijder commit",
		MoveDownCommit:                      "Verplaats commit 1 naar beneden",
		MoveUpCommit:                        "Verplaats commit 1 naar boven",
		EditCommit:                          "Wijzig commit",
		AmendToCommit:                       "Wijzig commit met staged veranderingen",
		RenameCommitEditor:                  "Hernoem commit met editor",
		NoCommitsThisBranch:                 "Geen commits in deze branch",
		Error:                               "Foutmelding",
		PickHunk:                            "Kies stuk",
		PickAllHunks:                        "Kies beide stukken",
		Undo:                                "Ongedaan maken",
		UndoReflog:                          "Ongedaan maken (via reflog) (experimenteel)",
		RedoReflog:                          "Redo (via reflog) (experimenteel)",
		Pop:                                 "Pop",
		Drop:                                "Laten vallen",
		Apply:                               "Toepassen",
		NoStashEntries:                      "Geen stash items",
		StashDrop:                           "Stash laten vallen",
		SureDropStashEntry:                  "Weet je het zeker dat je deze stash entry wil laten vallen?",
		StashPop:                            "Stash pop",
		SurePopStashEntry:                   "Weet je zeker dat je deze stash entry wil poppen?",
		StashApply:                          "Stash toepassen",
		SureApplyStashEntry:                 "Weet je zeker dat je deze stash entry wil toepassen?",
		NoTrackedStagedFilesStash:           "Je hebt geen tracked/staged bestanden om te laten stashen",
		StashChanges:                        "Stash veranderingen",
		RenameStash:                         "Rename stash",
		RenameStashPrompt:                   "Rename stash: {{.stashName}}",
		NoChangedFiles:                      "Geen veranderde bestanden",
		OpenConfig:                          "Open config bestand",
		EditConfig:                          "Verander config bestand",
		ForcePush:                           "Forceer push",
		ForcePushPrompt:                     "Jouw branch is afgeweken van de remote branch. Druk 'esc' om te annuleren, of 'enter' om geforceert te pushen.",
		ForcePushDisabled:                   "Your branch has diverged from the remote branch and you've disabled force pushing",
		UpdatesRejectedAndForcePushDisabled: "Updates were rejected and you have disabled force pushing",
		CheckForUpdate:                      "Check voor updates",
		CheckingForUpdates:                  "Zoeken naar updates...",
		OnLatestVersionErr:                  "Je hebt al de laatste versie",
		MajorVersionErr:                     "Nieuwe versie ({{.newVersion}}) is niet backwards compatibele vergeleken met de huidige versie ({{.currentVersion}})",
		CouldNotFindBinaryErr:               "Kon geen binary vinden op {{.url}}",
		IntroPopupMessage:                   "Bedankt voor het gebruik maken van lazygit! 2 dingen die je moet weten:\n\n1) Als je meer van lazygit zijn features wilt leren bekijk dan deze video:\n   https://youtu.be/CPLdltN7wgE\n\n2) Als je git gebruikt, ben je een programmeur! Met jouw hulp kunnen we lazygit verbeteren, dus overweeg om een ​​donateur te worden en mee te doen aan het plezier op\n   https://github.com/jesseduffield/lazygit",
		GitconfigParseErr:                   `Gogit kon je gitconfig bestand niet goed parsen door de aanwezigheid van losstaande '\' tekens. Het weghalen van deze tekens zou het probleem moeten oplossen. `,
		EditFile:                            `Verander bestand`,
		OpenFile:                            `Open bestand`,
		IgnoreFile:                          `Voeg toe aan .gitignore`,
		RefreshFiles:                        `Refresh bestanden`,
		MergeIntoCurrentBranch:              `Merge in met huidige checked out branch`,
		ConfirmQuit:                         `Weet je zeker dat je dit programma wil sluiten?`,
		SwitchRepo:                          "Wissel naar een recente repo",
		AllBranchesLogGraph:                 `Alle logs van de branch laten zien`,
		UnsupportedGitService:               `Niet-ondersteunde git-service`,
		CreatePullRequest:                   `Maak een pull-request`,
		CopyPullRequestURL:                  `Kopieer de URL van het pull-verzoek naar het klembord`,
		NoBranchOnRemote:                    `Deze branch bestaat niet op de remote. U moet het eerst naar de remote pushen.`,
		Fetch:                               `Fetch`,
		NoAutomaticGitFetchTitle:            `Geen automatische git fetch`,
		NoAutomaticGitFetchBody:             `Lazygit kan niet "git fetch" uitvoeren in een privé repository, gebruik f in het branches paneel om "git fetch" manueel uit te voeren`,
		FileEnter:                           `Stage individuele hunks/lijnen`,
		FileStagingRequirements:             `Kan alleen individuele lijnen stagen van getrackte bestanden met onstaged veranderingen`,
		StageSelection:                      `Toggle lijnen staged / unstaged`,
		ResetSelection:                      `Verwijdert change (git reset)`,
		ToggleDragSelect:                    `Toggle drag selecteer`,
		ToggleSelectHunk:                    `Toggle selecteer hunk`,
		ToggleSelectionForPatch:             `Voeg toe/verwijder lijn(en) in patch`,
		ToggleStagingPanel:                  `Ga naar een ander paneel`,
		ReturnToFilesPanel:                  `Ga terug naar het bestanden paneel`,
		FastForward:                         `Fast-forward deze branch vanaf zijn upstream`,
		Fetching:                            "Fetching en fast-forwarding {{.from}} -> {{.to}} ...",
		FoundConflictsTitle:                 "Conflicten!",
		ViewMergeRebaseOptions:              "Bekijk merge/rebase opties",
		NotMergingOrRebasing:                "Je bent momenteel niet aan het rebasen of mergen",
		RecentRepos:                         "Recente repositories",
		MergeOptionsTitle:                   "Merge opties",
		RebaseOptionsTitle:                  "Rebase opties",
		CommitMessageTitle:                  "Commit bericht",
		LocalBranchesTitle:                  "Branches",
		SearchTitle:                         "Zoek",
		TagsTitle:                           "Tags",
		MenuTitle:                           "Menu",
		RemotesTitle:                        "Remotes",
		RemoteBranchesTitle:                 "Remote branches",
		PatchBuildingTitle:                  "Patch bouwen",
		InformationTitle:                    "Informatie",
		SecondaryTitle:                      "Secondary",
		ReflogCommitsTitle:                  "Reflog",
		GlobalTitle:                         "Globale sneltoetsen",
		ConflictsResolved:                   "Alle merge conflicten zijn opgelost. Wilt je verder gaan?",
		MergingTitle:                        "Mergen",
		ConfirmMerge:                        "Weet je zeker dat je '{{.selectedBranch}}' in '{{.checkedOutBranch}}' wil mergen?",
		FwdNoUpstream:                       "Kan niet de branch vooruitspoelen zonder upstream",
		FwdCommitsToPush:                    "Je kan niet vooruitspoelen als de branch geen nieuwe commits heeft",
		ErrorOccurred:                       "Er is iets fout gegaan! Zou je hier een issue aan willen maken",
		NoRoom:                              "Niet genoeg ruimte",
		YouAreHere:                          "JE BENT HIER",
		RewordNotSupported:                  "Herformatteren van commits in interactief rebasen is nog niet ondersteund",
		CherryPickCopy:                      "Kopieer commit (cherry-pick)",
		CherryPickCopyRange:                 "Kopieer commit reeks (cherry-pick)",
		PasteCommits:                        "Plak commits (cherry-pick)",
		SureCherryPick:                      "Weet je zeker dat je de gekopieerde commits naar deze branch wil cherry-picken?",
		CherryPick:                          "Cherry-Pick",
		Donate:                              "Doneer",
		PrevLine:                            "Selecteer de vorige lijn",
		NextLine:                            "Selecteer de volgende lijn",
		PrevHunk:                            "Selecteer de vorige hunk",
		NextHunk:                            "Selecteer de volgende hunk",
		PrevConflict:                        "Selecteer voorgaand conflict",
		NextConflict:                        "Selecteer volgende conflict",
		SelectPrevHunk:                      "Selecteer bovenste hunk",
		SelectNextHunk:                      "Selecteer onderste hunk",
		ScrollDown:                          "Scroll omlaag",
		ScrollUp:                            "Scroll omhoog",
		ScrollUpMainPanel:                   "Scroll naar beneden vanaf hoofdpaneel",
		ScrollDownMainPanel:                 "Scroll naar beneden vanaf hoofdpaneel",
		AmendCommitTitle:                    "Commit wijzigen",
		AmendCommitPrompt:                   "Weet je zeker dat je deze commit wil wijzigen met de vorige staged bestanden?",
		DeleteCommitTitle:                   "Verwijder commit",
		DeleteCommitPrompt:                  "Weet je zeker dat je deze commit wil verwijderen?",
		SquashingStatus:                     "Squashen",
		FixingStatus:                        "Fixing up",
		DeletingStatus:                      "Verwijderen",
		MovingStatus:                        "Verplaatsen",
		RebasingStatus:                      "Rebasen",
		AmendingStatus:                      "Wijzigen",
		CherryPickingStatus:                 "Cherry-picken",
		UndoingStatus:                       "Ongedaan maken",
		RedoingStatus:                       "Redoing",
		CheckingOutStatus:                   "Uitchecken",
		CommitFiles:                         "Commit bestanden",
		ViewItemFiles:                       "Bekijk gecommite bestanden",
		CommitFilesTitle:                    "Commit bestanden",
		CheckoutCommitFile:                  "Bestand uitchecken",
		DiscardOldFileChange:                "Uitsluit deze commit zijn veranderingen aan dit bestand",
		DiscardFileChangesTitle:             "Uitsluit bestand zijn veranderingen",
		DiscardFileChangesPrompt:            "Weet je zeker dat je de wijzigingen van deze commit in dit bestand wilt weggooien? Als dit bestand is gecreëerd in deze commit dan zal dit bestand worden verwijdert",
		DisabledForGPG:                      "Onderdelen niet beschikbaar voor gebruikers die GPG gebruiken",
		CreateRepo:                          "Niet in een git repository. Creëer een nieuwe git repository? (y/n): ",
		AutoStashTitle:                      "Autostash?",
		AutoStashPrompt:                     "Je moet je veranderingen stashen en poppen om ze over te brengen. Dit automatisch doen? (enter/esc)",
		StashPrefix:                         "Auto-stashing veranderingen voor ",
		ViewDiscardOptions:                  "Bekijk 'veranderingen ongedaan maken' opties",
		Cancel:                              "Annuleren",
		DiscardAllChanges:                   "Negeer alle wijzigingen",
		DiscardUnstagedChanges:              "Negeer unstaged wijzigingen",
		DiscardAllChangesToAllFiles:         "Verwijder werkende tree",
		DiscardAnyUnstagedChanges:           "Gooi unstaged wijzigingen weg",
		DiscardUntrackedFiles:               "Negeer niet-gevonden bestanden",
		ViewResetOptions:                    `Bekijk reset opties`,
		HardReset:                           "Harde reset",
		CreateFixupCommit:                   `Creëer fixup commit voor deze commit`,
		SquashAboveCommits:                  `Squash bovenstaande commits`,
		SureSquashAboveCommits:              `Weet je zeker dat je alles wil squash/fixup! voor de bovenstaand commits {{.commit}}?`,
		CreateFixupCommitDescription:        `Creëer fixup commit`,
		SureCreateFixupCommit:               `Weet je zeker dat je een fixup wil maken! commit voor commit {{.commit}}?`,
		ExecuteCustomCommand:                "Voer aangepaste commando uit",
		CustomCommand:                       "Aangepaste commando:",
		CommitChangesWithoutHook:            "Commit veranderingen zonder pre-commit hook",
		SkipHookPrefixNotConfigured:         "Je hebt nog niet een commit bericht voorvoegsel ingesteld voor het overslaan van hooks. Set `git.skipHookPrefix = 'WIP'` in je config",
		ResetTo:                             `Reset naar`,
		PressEnterToReturn:                  "Press om terug te gaan naar lazygit",
		ViewStashOptions:                    "Bekijk stash opties",
		StashAllChanges:                     "Stash-bestanden",
		StashAllChangesKeepIndex:            "Stash staged wijzigingen",
		StashOptions:                        "Stash opties",
		NotARepository:                      "Fout: moet in een git repository uitgevoerd worden",
		Jump:                                "Ga naar paneel",
		DiscardPatch:                        "Patch weg gooien",
		DiscardPatchConfirm:                 "Je kan alleen maar een patch bouwen van 1 commit. Huidige patch weggooien?",
		CantPatchWhileRebasingError:         "Je kan geen patch bouwen of patch commando uitvoeren wanneer je in een merging of rebasing state zit",
		ToggleAddToPatch:                    "Toggle bestand inbegrepen in patch",
		ViewPatchOptions:                    "Bekijk aangepaste patch opties",
		PatchOptionsTitle:                   "Patch opties",
		NoPatchError:                        "Nog geen patch gecreëerd. Om een patch te bouwen gebruik 'space' op een commit bestand of 'enter' om een spesiefieke lijnen toe te voegen",
		EnterFile:                           "Enter bestand om geselecteerde regels toe te voegen aan de patch",
		ExitCustomPatchBuilder:              `Sluit lijn-bij-lijn modus`,
		EnterUpstream:                       `Enter upstream als '<remote> <branchnaam>'`,
		ReturnToRemotesList:                 `Ga terug naar remotes lijst`,
		AddNewRemote:                        `Voeg een nieuwe remote toe`,
		NewRemoteName:                       `Nieuwe remote name:`,
		NewRemoteUrl:                        `Nieuwe remote url:`,
		EditRemoteName:                      `Enter updated remote naam voor {{.remoteName}}:`,
		EditRemoteUrl:                       `Enter updated remote url voor {{.remoteName}}:`,
		RemoveRemote:                        `Verwijder remote`,
		RemoveRemotePrompt:                  "Weet je zeker dat je deze remote wilt verwijderen",
		DeleteRemoteBranch:                  "Verwijder remote branch",
		DeleteRemoteBranchMessage:           "Weet je zeker dat je deze remote branch wilt verwijderen",
		SetUpstream:                         "Stel in als upstream van uitgecheckte branch",
		SetAsUpstream:                       "Stel in als upstream van uitgecheckte branch",
		SetUpstreamTitle:                    "Stel in als upstream branch",
		SetUpstreamMessage:                  "Weet je zeker dat je de upstream branch van '{{.checkedOut}}' naar '{{.selected}}' wilt zetten",
		EditRemote:                          "Wijzig remote",
		TagCommit:                           "Tag commit",
		TagNameTitle:                        "Tag naam:",
		DeleteTag:                           "Verwijder tag",
		DeleteTagTitle:                      "Verwijder tag",
		DeleteTagPrompt:                     "Weet je zeker dat je '{{.tagName}}' wil verwijderen?",
		PushTagTitle:                        "Remote om tag '{{.tagName}}' te pushen naar:",
		PushTag:                             "Push tag",
		CreateTag:                           "Creëer tag",
		CreateTagTitle:                      "Tag naam:",
		FetchRemote:                         "Fetch remote",
		FetchingRemoteStatus:                "Remote fetchen",
		CheckoutCommit:                      "Checkout commit",
		SureCheckoutThisCommit:              "Weet je zeker dat je deze commit wil uitchecken?",
		GitFlowOptions:                      "Laat git-flow opties zien",
		NotAGitFlowBranch:                   "Dit lijkt geen git flow branch te zijn",
		NewGitFlowBranchPrompt:              "Nieuwe '{{.branchType}}' naam:",
		IgnoreTracked:                       "Negeer tracked bestand",
		IgnoreTrackedPrompt:                 "Weet je zeker dat je een getracked bestand wil negeren?",
		ViewResetToUpstreamOptions:          "Bekijk upstream reset opties",
		NextScreenMode:                      "Volgende scherm modus (normaal/half/groot)",
		PrevScreenMode:                      "Vorige scherm modus",
		StartSearch:                         "Start met zoeken",
		Panel:                               "Paneel",
		Keybindings:                         "Sneltoetsen",
		RenameBranch:                        "Hernoem branch",
		NewBranchNamePrompt:                 "Noem een nieuwe branch naam",
		RenameBranchWarning:                 "Deze branch volgt een remote. Deze actie zal alleen de locale branch name wijzigen niet de naam van de remote branch. Verder gaan?",
		OpenMenu:                            "Open menu",
		ResetCherryPick:                     "Reset cherry-picked (gekopieerde) commits selectie",
		NextTab:                             "Volgende tabblad",
		PrevTab:                             "Vorige tabblad",
		CantUndoWhileRebasing:               "Kan niet ongedaan maken terwijl je aan het rebasen bent",
		CantRedoWhileRebasing:               "Kan niet opnieuw doen (redo) terwijl je aan het rebasen bent",
		MustStashWarning:                    "Een patch in de index stoppen vereist stashen en onstashen van je wijzigingen. Als er iets verkeert gaat kan je je bestanden terug vinden in de stash. Verder gaan?",
		MustStashTitle:                      "Moet stashen",
		ConfirmationTitle:                   "Bevestigingspaneel",
		PrevPage:                            "Vorige pagina",
		NextPage:                            "Volgende pagina",
		GotoTop:                             "Scroll naar boven",
		GotoBottom:                          "Scroll naar beneden",
		FilteringBy:                         "Filteren bij",
		ResetInParentheses:                  "(reset)",
		OpenFilteringMenu:                   "Bekijk scoping opties",
		FilterBy:                            "Filter bij",
		ExitFilterMode:                      "Stop met filteren bij pad",
		FilterPathOption:                    "Vulin pad om op te filteren",
		EnterFileName:                       "Vulin path:",
		FilteringMenuTitle:                  "Filteren",
		MustExitFilterModeTitle:             "Command niet beschikbaar",
		MustExitFilterModePrompt:            "Command niet beschikbaar in filter modus. Sluit filter modus?",
		Diff:                                "Diff",
		EnterRefToDiff:                      "Vul in ref naar diff",
		EnteRefName:                         "Vul in ref:",
		ExitDiffMode:                        "Sluit diff mode",
		DiffingMenuTitle:                    "Diffen",
		SwapDiff:                            "Keer diff richting om",
		OpenDiffingMenu:                     "Open diff menu",
		ShowingGitDiff:                      "Laat output zien voor:",
		CopyCommitShaToClipboard:            "Kopieer commit SHA naar klembord",
		CopyCommitMessageToClipboard:        "Kopieer commit bericht naar klembord",
		CopyBranchNameToClipboard:           "Kopieer branch name naar klembord",
		CopyFileNameToClipboard:             "Kopieer de bestandsnaam naar het klembord",
		CopyCommitFileNameToClipboard:       "Kopieer de vastgelegde bestandsnaam naar het klembord",
		CommitPrefixPatternError:            "Fout in commitPrefix patroon",
		NoFilesStagedTitle:                  "Geen bestanden gestaged",
		NoFilesStagedPrompt:                 "Je hebt geen bestanden gestaged. Commit alle bestanden?",
		BranchNotFoundTitle:                 "Branch niet gevonden",
		BranchNotFoundPrompt:                "Branch niet gevonden. Creëer een nieuwe branch genaamd",
		PullRequestURLCopiedToClipboard:     "Pull-aanvraag-URL gekopieerd naar klembord",
		CommitMessageCopiedToClipboard:      "Commit message gekopieerd naar klembord",
		CopiedToClipboard:                   "Gekopieerd naar klembord",
		NavigationTitle:                     "Lijstpaneel navigatie",
		ViewCommits:                         "Bekijk commits",
		ToggleTreeView:                      "Toggle bestandsboom weergave",
		CreateNewBranchFromCommit:           "Creëer nieuwe branch van commit",
		CopySubmoduleNameToClipboard:        "Kopieer submodule naam naar klembord",
		EnterSubmodule:                      "Enter submodule",
		AddSubmodule:                        "Voeg nieuwe submodule toe",
		InitSubmodule:                       "Initialiseer submodule",
		ViewBulkSubmoduleOptions:            "Bekijk bulk submodule opties",
		CreatePullRequestOptions:            "Bekijk opties voor pull-aanvraag",
		ConfirmRevertCommit:                 "Weet u zeker dat u {{.selectedCommit}} ongedaan wilt maken?",
	}
}
